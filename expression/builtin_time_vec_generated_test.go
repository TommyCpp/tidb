// Copyright 2019 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go generate in expression/generator; DO NOT EDIT.

package expression

import (
	"testing"

	. "github.com/pingcap/check"
	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/mysql"
	"github.com/pingcap/tidb/types"
)

type gener struct {
	defaultGener
}

func (g gener) gen() interface{} {
	result := g.defaultGener.gen()
	if _, ok := result.(string); ok {
		dg := &defaultGener{eType: types.ETDuration, nullRation: 0}
		d := dg.gen().(types.Duration)
		if int8(d.Duration)%2 == 0 {
			d.Fsp = 0
		} else {
			d.Fsp = 1
		}
		result = d.String()
	}
	return result
}

var vecBuiltinTimeGeneratedCases = map[string][]vecExprBenchCase{
	ast.AddTime: {
		// builtinAddDatetimeAndDurationSig
		{
			retEvalType:   types.ETDatetime,
			childrenTypes: []types.EvalType{types.ETDatetime, types.ETDuration},
			geners: []dataGenerator{
				gener{defaultGener{eType: types.ETDatetime, nullRation: 0.2}},
				gener{defaultGener{eType: types.ETDuration, nullRation: 0.2}},
			},
		},
		// builtinAddDatetimeAndStringSig
		{
			retEvalType:   types.ETDatetime,
			childrenTypes: []types.EvalType{types.ETDatetime, types.ETString},
			geners: []dataGenerator{
				gener{defaultGener{eType: types.ETDatetime, nullRation: 0.2}},
				gener{defaultGener{eType: types.ETString, nullRation: 0.2}},
			},
		},
		// builtinAddDurationAndDurationSig
		{
			retEvalType:   types.ETDuration,
			childrenTypes: []types.EvalType{types.ETDuration, types.ETDuration},
			geners: []dataGenerator{
				gener{defaultGener{eType: types.ETDuration, nullRation: 0.2}},
				gener{defaultGener{eType: types.ETDuration, nullRation: 0.2}},
			},
		},
		// builtinAddDurationAndStringSig
		{
			retEvalType:   types.ETDuration,
			childrenTypes: []types.EvalType{types.ETDuration, types.ETString},
			geners: []dataGenerator{
				gener{defaultGener{eType: types.ETDuration, nullRation: 0.2}},
				gener{defaultGener{eType: types.ETString, nullRation: 0.2}},
			},
		},
		// builtinAddStringAndDurationSig
		{
			retEvalType:   types.ETString,
			childrenTypes: []types.EvalType{types.ETString, types.ETDuration},
			geners: []dataGenerator{
				gener{defaultGener{eType: types.ETString, nullRation: 0.2}},
				gener{defaultGener{eType: types.ETDuration, nullRation: 0.2}},
			},
		},
		// builtinAddStringAndStringSig
		{
			retEvalType:   types.ETString,
			childrenTypes: []types.EvalType{types.ETString, types.ETString},
			geners: []dataGenerator{
				gener{defaultGener{eType: types.ETString, nullRation: 0.2}},
				gener{defaultGener{eType: types.ETString, nullRation: 0.2}},
			},
		},
		// builtinAddDateAndDurationSig
		{
			retEvalType:        types.ETString,
			childrenTypes:      []types.EvalType{types.ETDuration, types.ETDuration},
			childrenFieldTypes: []*types.FieldType{types.NewFieldType(mysql.TypeDate), types.NewFieldType(mysql.TypeDuration)},
			geners: []dataGenerator{
				gener{defaultGener{eType: types.ETDuration, nullRation: 0.2}},
				gener{defaultGener{eType: types.ETDuration, nullRation: 0.2}},
			},
		},
		// builtinAddDateAndStringSig
		{
			retEvalType:        types.ETString,
			childrenTypes:      []types.EvalType{types.ETDuration, types.ETString},
			childrenFieldTypes: []*types.FieldType{types.NewFieldType(mysql.TypeDate), types.NewFieldType(mysql.TypeString)},
			geners: []dataGenerator{
				gener{defaultGener{eType: types.ETDuration, nullRation: 0.2}},
				gener{defaultGener{eType: types.ETString, nullRation: 0.2}},
			},
		},
		// builtinAddTimeDateTimeNullSig
		{
			retEvalType:   types.ETDatetime,
			childrenTypes: []types.EvalType{types.ETDatetime, types.ETDatetime},
			geners: []dataGenerator{
				gener{defaultGener{eType: types.ETDatetime, nullRation: 0.2}},
				gener{defaultGener{eType: types.ETDatetime, nullRation: 0.2}},
			},
		},
		// builtinAddTimeStringNullSig
		{
			retEvalType:        types.ETString,
			childrenTypes:      []types.EvalType{types.ETDatetime, types.ETDatetime},
			childrenFieldTypes: []*types.FieldType{types.NewFieldType(mysql.TypeDate), types.NewFieldType(mysql.TypeDatetime)},
			geners: []dataGenerator{
				gener{defaultGener{eType: types.ETDatetime, nullRation: 0.2}},
				gener{defaultGener{eType: types.ETDatetime, nullRation: 0.2}},
			},
		},
		// builtinAddTimeDurationNullSig
		{
			retEvalType:   types.ETDuration,
			childrenTypes: []types.EvalType{types.ETDuration, types.ETDatetime},
			geners: []dataGenerator{
				gener{defaultGener{eType: types.ETDuration, nullRation: 0.2}},
				gener{defaultGener{eType: types.ETDatetime, nullRation: 0.2}},
			},
		},
	},

	ast.TimeDiff: {
		// builtinNullTimeDiffSig
		{retEvalType: types.ETDuration, childrenTypes: []types.EvalType{types.ETDuration, types.ETDatetime}},
		{retEvalType: types.ETDuration, childrenTypes: []types.EvalType{types.ETDuration, types.ETTimestamp}},
		{retEvalType: types.ETDuration, childrenTypes: []types.EvalType{types.ETDatetime, types.ETDuration}},
		{retEvalType: types.ETDuration, childrenTypes: []types.EvalType{types.ETTimestamp, types.ETDuration}},
		// builtinDurationDurationTimeDiffSig
		{retEvalType: types.ETDuration, childrenTypes: []types.EvalType{types.ETDuration, types.ETDuration}},
		// builtinDurationStringTimeDiffSig
		{retEvalType: types.ETDuration, childrenTypes: []types.EvalType{types.ETDuration, types.ETString}, geners: []dataGenerator{nil, &timeStrGener{Year: 2019, Month: 10}}},
		{retEvalType: types.ETDuration, childrenTypes: []types.EvalType{types.ETDuration, types.ETString}, geners: []dataGenerator{nil, &dateTimeStrGener{Year: 2019, Month: 10}}},
		{retEvalType: types.ETDuration, childrenTypes: []types.EvalType{types.ETDuration, types.ETString}, geners: []dataGenerator{nil, &dateTimeStrGener{Year: 2019, Month: 10, Fsp: 4}}},
		// builtinTimeTimeTimeDiffSig
		{retEvalType: types.ETDuration, childrenTypes: []types.EvalType{types.ETDatetime, types.ETDatetime}, geners: []dataGenerator{&dateTimeGener{Year: 2019, Month: 10}, &dateTimeGener{Year: 2019, Month: 10}}},
		{retEvalType: types.ETDuration, childrenTypes: []types.EvalType{types.ETDatetime, types.ETTimestamp}, geners: []dataGenerator{&dateTimeGener{Year: 2019, Month: 10}, &dateTimeGener{Year: 2019, Month: 10}}},
		{retEvalType: types.ETDuration, childrenTypes: []types.EvalType{types.ETTimestamp, types.ETTimestamp}, geners: []dataGenerator{&dateTimeGener{Year: 2019, Month: 10}, &dateTimeGener{Year: 2019, Month: 10}}},
		{retEvalType: types.ETDuration, childrenTypes: []types.EvalType{types.ETTimestamp, types.ETDatetime}, geners: []dataGenerator{&dateTimeGener{Year: 2019, Month: 10}, &dateTimeGener{Year: 2019, Month: 10}}},
		// builtinTimeStringTimeDiffSig
		{retEvalType: types.ETDuration, childrenTypes: []types.EvalType{types.ETDatetime, types.ETString}, geners: []dataGenerator{&dateTimeGener{Year: 2019, Month: 10}, &timeStrGener{Year: 2019, Month: 10}}},
		{retEvalType: types.ETDuration, childrenTypes: []types.EvalType{types.ETDatetime, types.ETString}, geners: []dataGenerator{&dateTimeGener{Year: 2019, Month: 10}, &dateTimeStrGener{Year: 2019, Month: 10}}},
		{retEvalType: types.ETDuration, childrenTypes: []types.EvalType{types.ETTimestamp, types.ETString}, geners: []dataGenerator{&dateTimeGener{Year: 2019, Month: 10}, &timeStrGener{Year: 2019, Month: 10}}},
		{retEvalType: types.ETDuration, childrenTypes: []types.EvalType{types.ETTimestamp, types.ETString}, geners: []dataGenerator{&dateTimeGener{Year: 2019, Month: 10}, &dateTimeStrGener{Year: 2019, Month: 10}}},
		// builtinStringDurationTimeDiffSig
		{retEvalType: types.ETDuration, childrenTypes: []types.EvalType{types.ETString, types.ETDuration}, geners: []dataGenerator{&timeStrGener{Year: 2019, Month: 10}, nil}},
		{retEvalType: types.ETDuration, childrenTypes: []types.EvalType{types.ETString, types.ETDuration}, geners: []dataGenerator{&dateTimeStrGener{Year: 2019, Month: 10}, nil}},
		// builtinStringTimeTimeDiffSig
		{retEvalType: types.ETDuration, childrenTypes: []types.EvalType{types.ETString, types.ETDatetime}, geners: []dataGenerator{&timeStrGener{Year: 2019, Month: 10}, &dateTimeGener{Year: 2019, Month: 10}}},
		{retEvalType: types.ETDuration, childrenTypes: []types.EvalType{types.ETString, types.ETDatetime}, geners: []dataGenerator{&dateTimeStrGener{Year: 2019, Month: 10}, &dateTimeGener{Year: 2019, Month: 10}}},
		{retEvalType: types.ETDuration, childrenTypes: []types.EvalType{types.ETString, types.ETTimestamp}, geners: []dataGenerator{&timeStrGener{Year: 2019, Month: 10}, &dateTimeGener{Year: 2019, Month: 10}}},
		{retEvalType: types.ETDuration, childrenTypes: []types.EvalType{types.ETString, types.ETTimestamp}, geners: []dataGenerator{&dateTimeStrGener{Year: 2019, Month: 10}, &dateTimeGener{Year: 2019, Month: 10}}},
		// builtinStringStringTimeDiffSig
		{retEvalType: types.ETDuration, childrenTypes: []types.EvalType{types.ETString, types.ETString}, geners: []dataGenerator{&timeStrGener{Year: 2019, Month: 10}, &dateTimeStrGener{Year: 2019, Month: 10}}},
		{retEvalType: types.ETDuration, childrenTypes: []types.EvalType{types.ETString, types.ETString}, geners: []dataGenerator{&dateTimeStrGener{Year: 2019, Month: 10}, &timeStrGener{Year: 2019, Month: 10}}},
	},
}

func (s *testEvaluatorSuite) TestVectorizedBuiltinTimeEvalOneVecGenerated(c *C) {
	testVectorizedEvalOneVec(c, vecBuiltinTimeGeneratedCases)
}

func (s *testEvaluatorSuite) TestVectorizedBuiltinTimeFuncGenerated(c *C) {
	testVectorizedBuiltinFunc(c, vecBuiltinTimeGeneratedCases)
}

func BenchmarkVectorizedBuiltinTimeEvalOneVecGenerated(b *testing.B) {
	benchmarkVectorizedEvalOneVec(b, vecBuiltinTimeGeneratedCases)
}

func BenchmarkVectorizedBuiltinTimeFuncGenerated(b *testing.B) {
	benchmarkVectorizedBuiltinFunc(b, vecBuiltinTimeGeneratedCases)
}
