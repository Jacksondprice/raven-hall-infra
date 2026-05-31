package handlers

import (
	"net/http/httptest"
	"testing"
)

func TestListInvoices(t *testing.T) {
	rec := httptest.NewRecorder()
	ListInvoices(rec, httptest.NewRequest("GET", "/api/invoices", nil))
	if rec.Code != 200 {
		t.Fatalf("status %d", rec.Code)
	}
}
