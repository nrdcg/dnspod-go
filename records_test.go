package dnspod

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestRecordsService_ListRecords_all(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/Record.List", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "unsupported method", http.StatusBadRequest)
			return
		}

		_, _ = fmt.Fprint(w, `{
			"status": {"code":"1","message":""},
			"records":[
				{"id":"44146112", "name":"yizerowwwww"},
				{"id":"44146112", "name":"yizerowwwww"}
			]}`)
	})

	records, _, err := client.Records.List("example.com", "")
	if err != nil {
		t.Fatal(err)
	}

	want := []Record{{ID: "44146112", Name: "yizerowwwww"}, {ID: "44146112", Name: "yizerowwwww"}}
	if !reflect.DeepEqual(records, want) {
		t.Errorf("got %+v, want %+v", records, want)
	}
}

func TestRecordsService_ListRecords_subdomain(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/Record.List", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "unsupported method", http.StatusBadRequest)
			return
		}

		_, _ = fmt.Fprint(w, `{
			"status": {"code":"1","message":""},
			"records":[
				{"id":"44146112", "name":"yizerowwwww"},
				{"id":"44146112", "name":"yizerowwwww"}
			]}`)
	})

	records, _, err := client.Records.List("11223344", "@")
	if err != nil {
		t.Fatal(err)
	}

	want := []Record{{ID: "44146112", Name: "yizerowwwww"}, {ID: "44146112", Name: "yizerowwwww"}}
	if !reflect.DeepEqual(records, want) {
		t.Errorf("got returned %+v, want %+v", records, want)
	}
}

func TestRecordsService_CreateRecord(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/Record.Create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "unsupported method", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		_, _ = fmt.Fprintf(w, `{"status": {"code":"1","message":""},"record":{"id":"26954449", "name":"@", "status":"enable"}}`)
	})

	recordValues := Record{Name: "@", Status: "enable"}
	record, _, err := client.Records.Create("44146112", recordValues)
	if err != nil {
		t.Fatal(err)
	}

	want := Record{ID: "26954449", Name: "@", Status: "enable"}
	if !reflect.DeepEqual(record, want) {
		t.Errorf("got %+v, want %+v", record, want)
	}
}

func TestRecordsService_GetRecord(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/Record.Info", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "unsupported method", http.StatusBadRequest)
			return
		}

		_, _ = fmt.Fprintf(w, `{"status": {"code":"1","message":""},"record":{"id":"26954449", "name":"@", "status":"enable"}}`)
	})

	record, _, err := client.Records.Get("44146112", "26954449")
	if err != nil {
		t.Fatal(err)
	}

	want := Record{ID: "26954449", Name: "@", Status: "enable"}
	if !reflect.DeepEqual(record, want) {
		t.Fatalf("got %+v, want %+v", record, want)
	}
}

func TestRecordsService_UpdateRecord(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/Record.Modify", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "unsupported method", http.StatusBadRequest)
			return
		}

		_, _ = fmt.Fprint(w, `{"status": {"code":"1","message":""},"record":{"id":26954449, "name":"@", "status":"enable"}}`)
	})

	recordValues := Record{ID: "26954449", Name: "@", Status: "enable"}
	record, _, err := client.Records.Update("44146112", "26954449", recordValues)
	if err != nil {
		t.Fatal(err)
	}

	want := RecordModify{ID: "26954449", Name: "@", Status: "enable"}
	if !reflect.DeepEqual(record, want) {
		t.Errorf("got %+v, want %+v", record, want)
	}
}

func TestRecordsService_DeleteRecord(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/Record.Remove", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "unsupported method", http.StatusBadRequest)
			return
		}

		_, _ = fmt.Fprint(w, `{"status": {"code":"1","message":""}}`)
	})

	_, err := client.Records.Delete("44146112", "26954449")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRecordsService_DeleteRecord_failed(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/Record.Remove", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "unsupported method", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(w, `{"message":"InvalID request"}`)
	})

	_, err := client.Records.Delete("44146112", "26954449")
	if err == nil {
		t.Fatal(err)
	}

	if match := "400 InvalID request"; !strings.Contains(err.Error(), match) {
		t.Errorf("got %+v, should match %+v", err, match)
	}
}
