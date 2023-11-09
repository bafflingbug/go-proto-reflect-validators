package validator

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
	"google.golang.org/protobuf/types/descriptorpb"
	"log"
	"regexp"
	"sync"
)

// regCache regexp cache
type regCache struct {
	sync.Map
}

// reset cache
func (r *regCache) reset() {
	r.Map = sync.Map{}
}

// Get get regexp instance
func (r *regCache) Get(expr string) (*regexp.Regexp, error) {
	if x, ok := r.Map.Load(expr); ok {
		if exp, ok := x.(*regexp.Regexp); ok {
			return exp, nil
		}
	}
	exp, err := regexp.Compile(expr)
	if err != nil {
		return nil, err
	}
	r.Map.Store(expr, exp)
	return exp, nil
}

var r = regCache{}

// ResetRegCache reset regexp cache
func ResetRegCache() {
	r.reset()
}

// validator proto validator
type validator struct {
	msg *dynamic.Message
}

// ValidMsg verify whether a proto message is legal
func ValidMsg(msg *dynamic.Message) (err error) {
	defer func() {
		if p := recover(); p != nil {
			log.Printf("[pb valid]panic: %s, msg: %+v", p, msg)
			err = nil
		}
	}()
	v := validator{
		msg: msg,
	}
	return v.Valid()
}

// Valid valid proto msg
func (v *validator) Valid() error {
	if v.msg == nil {
		return nil
	}
	fields := v.msg.GetKnownFields()
	for _, field := range fields {
		if field.IsExtension() {
			continue
		}

		value, err := v.msg.TryGetField(field)
		if err != nil {
			log.Printf("[pb valid]get field[%+v] value err: %s", field, err)
			continue
		}
		rule := v.getRule(field)

		if field.IsMap() {
			if err2 := v.validMap(field, value, rule); err2 != nil {
				return err2
			}
		} else if field.IsRepeated() {
			if err2 := v.validRepeated(field, value, rule); err2 != nil {
				return err2
			}
		} else {
			if err2 := v.validField(field, value, rule); err2 != nil {
				return err2
			}
		}

		//fmt.Println(field)
	}
	return nil
}

// getRule get verification rules
func (v *validator) getRule(field *desc.FieldDescriptor) *FieldValidator {
	opt := field.GetFieldOptions()
	if opt == nil {
		return nil
	}
	ext, err := proto.GetExtension(opt, E_Field)
	if err != nil {
		return nil
	}
	rule, ok := ext.(*FieldValidator)
	if !ok {
		return nil
	}
	return rule
}

// validRepeated valid list
func (v *validator) validRepeated(field *desc.FieldDescriptor, value interface{}, rule *FieldValidator) error {
	if value == nil {
		return nil
	}
	vList, ok := value.([]interface{})
	if !ok {
		log.Printf("[pb valid]field[%+v] value[%+v] is not array", field, value)
		return nil
	}

	if err := v.checkRepeated(field, vList, rule); err != nil {
		return err
	}

	for _, item := range vList {
		if err := v.validField(field, item, rule); err != nil {
			return err
		}
	}
	return nil
}

// validMap valid map
func (v *validator) validMap(field *desc.FieldDescriptor, value interface{}, rule *FieldValidator) error {
	if value == nil {
		return nil
	}
	vList, ok := value.(map[interface{}]interface{})
	if !ok {
		log.Printf("[pb valid]field[%+v] value[%+v] is not map", field, value)
		return nil
	}

	for key, item := range vList {
		if err := v.validField(field.GetMapKeyType(), key, rule); err != nil {
			return err
		}

		if err := v.validField(field.GetMapValueType(), item, nil); err != nil {
			return err
		}
	}
	return nil
}

// validField valid a field
func (v *validator) validField(field *desc.FieldDescriptor, value interface{}, rule *FieldValidator) error {
	if value == nil {
		return nil
	}

	switch field.GetType() {
	case descriptorpb.FieldDescriptorProto_TYPE_MESSAGE:
		//message
		return v.checkMessage(field, value, rule)

	case descriptorpb.FieldDescriptorProto_TYPE_INT32,
		descriptorpb.FieldDescriptorProto_TYPE_SINT32,
		descriptorpb.FieldDescriptorProto_TYPE_SFIXED32:
		//int32
		return v.checkInt(field, int64(value.(int32)), rule)

	case descriptorpb.FieldDescriptorProto_TYPE_INT64,
		descriptorpb.FieldDescriptorProto_TYPE_SINT64,
		descriptorpb.FieldDescriptorProto_TYPE_SFIXED64:
		//int64
		return v.checkInt(field, value.(int64), rule)

	case descriptorpb.FieldDescriptorProto_TYPE_UINT32,
		descriptorpb.FieldDescriptorProto_TYPE_FIXED32:
		//uint32
		return v.checkInt(field, int64(value.(uint32)), rule)

	case descriptorpb.FieldDescriptorProto_TYPE_UINT64,
		descriptorpb.FieldDescriptorProto_TYPE_FIXED64:
		//uint64
		return v.checkInt(field, int64(value.(uint64)), rule)

	case descriptorpb.FieldDescriptorProto_TYPE_FLOAT:
		//float32
		return v.checkFloat(field, float64(value.(float32)), rule)

	case descriptorpb.FieldDescriptorProto_TYPE_DOUBLE:
		//float64
		return v.checkFloat(field, value.(float64), rule)

	case descriptorpb.FieldDescriptorProto_TYPE_STRING:
		//string
		return v.checkString(field, value.(string), rule)

	case descriptorpb.FieldDescriptorProto_TYPE_BYTES:
		//[]bytes
		return v.checkBytes(field, value.([]byte), rule)

	case descriptorpb.FieldDescriptorProto_TYPE_ENUM:
		//enum
		return v.checkEnum(field, value.(int32), rule)
	}
	return nil
}

// checkRepeated check list
func (v *validator) checkRepeated(field *desc.FieldDescriptor, values []interface{}, rule *FieldValidator) error {
	if rule == nil {
		return nil
	}

	_len := int64(len(values))
	if rule.RepeatedCountMin != nil && !(_len >= *rule.RepeatedCountMin) {
		return ValidFail(field, "RepeatedCountMin", *rule.RepeatedCountMin, _len)
	}
	if rule.RepeatedCountMax != nil && !(_len <= *rule.RepeatedCountMax) {
		return ValidFail(field, "RepeatedCountMax", *rule.RepeatedCountMax, _len)
	}
	return nil
}

// checkMessage 检查消息
func (v *validator) checkMessage(field *desc.FieldDescriptor, value interface{}, rule *FieldValidator) error {
	subMsg, ok := value.(*dynamic.Message)
	if !ok {
		log.Printf("[pb valid]field[%+v] value[%+v] is not *dynamic.Message", field, value)
		return nil
	}
	if err := ValidMsg(subMsg); err != nil {
		return err
	}
	return nil
}

// checkInt check int
func (v *validator) checkInt(field *desc.FieldDescriptor, value int64, rule *FieldValidator) error {
	if rule == nil {
		return nil
	}

	if rule.IntGt != nil && !(value > *rule.IntGt) {
		return ValidFail(field, "IntGt", *rule.IntGt, value)
	}
	if rule.IntLt != nil && !(value < *rule.IntLt) {
		return ValidFail(field, "IntLt", *rule.IntLt, value)
	}
	return nil
}

// checkFloat check float
func (v *validator) checkFloat(field *desc.FieldDescriptor, value float64, rule *FieldValidator) error {
	if rule == nil {
		return nil
	}

	valueMax, valueMin := value, value
	if rule.FloatEpsilon != nil {
		//进行精度忽略
		valueMax += *rule.FloatEpsilon
		valueMin -= *rule.FloatEpsilon
	}

	if rule.FloatGt != nil && !(valueMax > *rule.FloatGt) {
		return ValidFail(field, "FloatGt", *rule.FloatGt, value)
	}
	if rule.FloatLt != nil && !(valueMin < *rule.FloatLt) {
		return ValidFail(field, "FloatLt", *rule.FloatLt, value)
	}

	if rule.FloatGte != nil && !(valueMax >= *rule.FloatGte) {
		return ValidFail(field, "FloatGte", *rule.FloatGte, value)
	}
	if rule.FloatLte != nil && !(valueMin <= *rule.FloatLte) {
		return ValidFail(field, "FloatLte", *rule.FloatLte, value)
	}
	return nil
}

// checkString check string
func (v *validator) checkString(field *desc.FieldDescriptor, value string, rule *FieldValidator) error {
	if rule == nil {
		return nil
	}

	if rule.StringNotEmpty != nil && *rule.StringNotEmpty && value == "" {
		return ValidFail(field, "StringNotEmpty", *rule.StringNotEmpty, value)
	}

	_len := int64(len(value))
	if rule.LengthGt != nil && !(_len > *rule.LengthGt) {
		return ValidFail(field, "LengthGt", *rule.LengthGt, _len)
	}
	if rule.LengthLt != nil && !(_len < *rule.LengthLt) {
		return ValidFail(field, "LengthLt", *rule.LengthLt, _len)
	}
	if rule.LengthEq != nil && !(_len == *rule.LengthEq) {
		return ValidFail(field, "LengthEq", *rule.LengthEq, _len)
	}

	if rule.Regex != nil {
		exp, err := r.Get(*rule.Regex)
		if err != nil {
			log.Printf("[pb valid]make regex[%s] err: %s", *rule.Regex, err)
		} else if !exp.MatchString(value) {
			return ValidFail(field, "Regex", *rule.Regex, value)
		}
	}

	return nil
}

// checkBytes check []byte
func (v *validator) checkBytes(field *desc.FieldDescriptor, value []byte, rule *FieldValidator) error {
	if rule == nil {
		return nil
	}

	_len := int64(len(value))
	if rule.LengthGt != nil && !(_len > *rule.LengthGt) {
		return ValidFail(field, "LengthGt", *rule.LengthGt, _len)
	}
	if rule.LengthLt != nil && !(_len < *rule.LengthLt) {
		return ValidFail(field, "LengthLt", *rule.LengthLt, _len)
	}
	if rule.LengthEq != nil && !(_len == *rule.LengthEq) {
		return ValidFail(field, "LengthEq", *rule.LengthEq, _len)
	}

	return nil
}

// checkEnum check enum
func (v *validator) checkEnum(field *desc.FieldDescriptor, value int32, rule *FieldValidator) error {
	if rule == nil || rule.IsInEnum == nil || !*rule.IsInEnum {
		return nil
	}

	for _, item := range field.GetEnumType().GetValues() {
		if value == item.GetNumber() {
			return nil
		}
	}
	return ValidFail(field, "IsInEnum", *rule.IsInEnum, false)
}

// ValidError error warp
type ValidError struct {
	field      *desc.FieldDescriptor
	validKey   string
	validValue interface{}
	fieldValue interface{}
}

// ValidFail error warp
func ValidFail(field *desc.FieldDescriptor, validKey string, validValue interface{}, fieldValue interface{}) error {
	return &ValidError{
		field:      field,
		validKey:   validKey,
		validValue: validValue,
		fieldValue: fieldValue,
	}
}

// Error implement interface
func (e *ValidError) Error() string {
	return fmt.Sprintf("[proto valid]error: field[%s (type:%s)] valid[%s(rule:%+v)] find[%+v]",
		e.field.GetName(), e.field.GetType(), e.validKey, e.validValue, e.fieldValue)
}
