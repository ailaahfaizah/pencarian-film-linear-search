package web

import (
    "html/template"
    "net/http"
    "strings"
    "time"

    "pencarian_film/data"
    "pencarian_film/search"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("web/templates/index.html")
    if err != nil {
        http.Error(w, "Template index error: "+err.Error(), 500)
        return
    }
    t.Execute(w, nil)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    query := strings.TrimSpace(r.FormValue("query"))

    // Jumlah pengulangan untuk menghindari waktu = 0 µs
    loops := 100000

    // --- Iterative ---
    start := time.Now()
    iterIndex := search.LinearIterative(data.Movies, query) // hitung index sekali saja
    for i := 0; i < loops; i++ {
        search.LinearIterative(data.Movies, query)
    }
    iterTime :=  float64( time.Since(start).Milliseconds()) 

    // --- Recursive ---
    start = time.Now()
    recIndex := search.LinearRecursive(data.Movies, query, 0)
    for i := 0; i < loops; i++ {
        search.LinearRecursive(data.Movies, query, 0)
    }
    recTime := float64( time.Since(start).Milliseconds()) 


    dataMap := map[string]interface{}{
        "Query":    query,
        "IterIdx":  iterIndex,
        "RecIdx":   recIndex,
        "IterTime": iterTime,
        "RecTime":  recTime,
    }

    t, err := template.ParseFiles("web/templates/result.html")
    if err != nil {
        http.Error(w, "Template result error: "+err.Error(), 500)
        return
    }

    t.Execute(w, dataMap)
}
