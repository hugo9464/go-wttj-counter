package main

import (
	"reflect"
	"testing"
)

func TestAppendIfMissing(t *testing.T) {
	tests := []struct {
		slice []string
		valueToAdd string
		expected []string
	}{
		{[]string{"a", "b", "c", "d"},"e", []string{"a", "b", "c", "d", "e"}},
		{[]string{"a", "b", "c", "d"},"c", []string{"a", "b", "c", "d"}},
	}
	for _, test := range tests {
		result := AppendIfMissing(test.slice, test.valueToAdd)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Result of AppendIfMissing incorrect with slice=%v and valueToAdd=%v. Expected: %v but got: %v", test.slice, test.valueToAdd, test.expected, result)
		}
	}
}

func TestAppendValue(t *testing.T) {
	tests := []struct {
		values   []Value
		category string
		expected []Value
	}{
		{[]Value{{category:"category1", count:1}, {category:"category2", count:0}}, "category2", []Value{{category:"category1", count:1}, {category:"category2", count:1}}},
		{[]Value{{category:"category1", count:1}}, "category2", []Value{{category:"category1", count:1}, {category:"category2", count:1}}},
	}
	for _, test := range tests {
		result := AppendValue(test.values, test.category)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Result of AppendValue incorrect with values=%v and category=%v. Expected: %v but got: %v", test.values, test.category, test.expected, result)
		}
	}
}

func TestCountDigits(t *testing.T) {
	tests := []struct {
		i int
		expected int
	}{
		{123, 3},
		{0, 1},
	}
	for _, test := range tests {
		result := CountDigits(test.i)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Result of CountDigits incorrect with i=%d Expected: %d but got: %d", test.i, test.expected, result)
		}
	}
}

func TestCreateNewValues(t *testing.T) {
	tests := []struct {
		category string
		expected []Value
	}{
		{"categoryName", []Value{{category:"categoryName", count:1}}},
	}
	for _, test := range tests {
		result := CreateNewValues(test.category)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Result of CreateNewValyes incorrect with category=%v Expected: %v but got: %v", test.category, test.expected, result)
		}
	}
}

func TestGetCategory(t *testing.T) {
	tests := []struct {
		professions [][]string
		id int
		expected string
	}{
		{[][]string {{"id", "name", "caregory_name"}, {"17", "INTERNSHIP", "Tech"}}, 17, "Tech"},
		{[][]string {{"id", "name", "caregory_name"}, {"17", "INTERNSHIP", "Tech"}}, 1, "Other"},
	}
	for _, test := range tests {
		result := GetCategory(test.professions, test.id)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Result of GetCategory incorrect with professions=%v and id=%d Expected: %v but got: %v", test.professions, test.id, test.expected, result)
		}
	}
}

func TestGetFirstRow(t *testing.T) {
	tests := []struct {
		categories []string
		longestContractSize int
		expected string
	}{
		{[]string{"category1", "category2", "category3"}, 10, "|          | category1 | category2 | category3 |"},
	}
	for _, test := range tests {
		result := GetFirstRow(test.categories, test.longestContractSize)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Result of GetFirstRow incorrect with categories=%v and longestContractSize=%d Expected: %v but got: %v", test.categories, test.longestContractSize, test.expected, result)
		}
	}
}

func TestGetStringOfChar(t *testing.T) {
	tests := []struct {
		s      string
		length int
		expected string
	}{
		{"-", 4, "----"},
		{"-", 0, ""},
	}
	for _, test := range tests {
		result := GetStringOfChar(test.s, test.length)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Result of GetStringOfChar incorrect with s=%v and length=%d Expected: %v but got: %v", test.s, test.length, test.expected, result)
		}
	}
}

func TestIncrementValue(t *testing.T) {
	tests := []struct {
		outputMap map[string][]Value
		contract  string
		category  string
		expected map[string][]Value
	}{
		{
			map[string][]Value{"contract1": {{category:"category1",count:1}, {category:"category2", count:1}}},
			"contract1" ,
			"category1",
			map[string][]Value{"contract1": {{category:"category1", count:2}, {category:"category2", count:1}}}},
		{
			map[string][]Value{"contract1": {{category:"category1",count:1}, {category:"category2", count:1}}},
			"contract2" ,
			"category1",
			map[string][]Value{"contract1": {{category:"category1", count:1}, {category:"category2", count:1}}, "contract2": {{category:"category1", count:1}}}},
		{
			map[string][]Value{"contract1": {{category:"category1",count:1}, {category:"category2", count:1}}},
			"contract1" ,
			"category3",
			map[string][]Value{"contract1": {{category:"category1", count:1}, {category:"category2", count:1}, {category:"category3", count:1}}}},
	}
	for _, test := range tests {
		result := IncrementValue(test.outputMap, test.contract, test.category)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Result of IncrementValue incorrect with outputMap=%v, contract=%v, category=%v Expected: %v but got: %v", test.outputMap, test.contract, test.category, test.expected, result)
		}
	}
}

func TestLongestStringSize(t *testing.T) {
	tests := []struct {
		strings []string
		expected int
	}{
		{
			[]string{"string1", "string12", "string123"},
			9},
	}
	for _, test := range tests {
		result := LongestStringSize(test.strings)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Result of LongestStringSize incorrect with strings=%v Expected: %d but got: %d", test.strings, test.expected, result)
		}
	}
}

func TestGetContractValues(t *testing.T) {
	tests := []struct {
		values     []Value
		categories []string
		expected string
	}{
		{
			[]Value{{category:"category1",count:1},{category:"category2",count:2},{category:"category3",count:3}},
			[]string{"category3", "category2", "category1"},
			"     3     |     2     |     1     |\n"},
	}
	for _, test := range tests {
		result := GetContractValues(test.values, test.categories)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Result of GetContractValues incorrect with values=%v and categories=%v Expected: '%v' but got: '%v'", test.values, test.categories, test.expected, result)
		}
	}
}

func TestGetValueForCategory(t *testing.T) {
	tests := []struct {
		values   []Value
		category string
		expected string
	}{
		{
			[]Value{{category:"category1",count:1},{category:"category2",count:2},{category:"category3",count:3}},
			"category2",
			"     2     |",},
	}
	for _, test := range tests {
		result := GetValueForCategory(test.values, test.category)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Result of GetValueForCategory incorrect with values=%v and category=%v Expected: '%v' but got: '%v'", test.values, test.category, test.expected, result)
		}
	}
}

func TestReadFile(t *testing.T) {
	tests := []struct {
		name string
		expected [][]string
	}{
		{
			"technical-test-jobs-sample.csv",
			[][]string{
				{"profession_id", "contract_type", "name", "office_latitude", "office_longitude"},
				{"1", "contract1", "name1", "1234", "5678"},
				{"2", "contract2", "name2", "5678", "1234"}}},
	}
	for _, test := range tests {
		result := ReadFile(test.name)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Result of ReadFile incorrect with name=%v Expected: '%v' but got: '%v'", test.name, test.expected, result)
		}
	}
}