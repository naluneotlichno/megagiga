package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test1(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=1000&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body.String())
}

func Test2(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=1000&city=fucking_moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	expected := "wrong city value"
	require.Equal(t, expected, responseRecorder.Body.String())
}

func Test3(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=1000&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	expected := "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент"
	require.Equal(t, expected, responseRecorder.Body.String())
}
