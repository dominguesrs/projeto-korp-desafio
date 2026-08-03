package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Response representa o JSON retornado pelo endpoint /projeto-korp
type Response struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

var (
	// requestsTotal conta o volume de requisicoes recebidas, segmentado por status
	requestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_server_projeto_korp_requests_total",
			Help: "Numero total de requisicoes recebidas pelo endpoint /projeto-korp",
		},
		[]string{"status"},
	)

	// serviceUp indica a disponibilidade do servico (1 = disponivel, 0 = indisponivel)
	serviceUp = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "http_server_projeto_korp_up",
			Help: "Indica se o servico http-server-projeto-korp esta disponivel",
		},
	)
)

// projetoKorpHandler atende ao endpoint GET /projeto-korp
func projetoKorpHandler(w http.ResponseWriter, r *http.Request) {
	resp := Response{
		Nome:    "Projeto Korp",
		Horario: time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		requestsTotal.WithLabelValues("error").Inc()
		http.Error(w, "erro interno ao gerar resposta", http.StatusInternalServerError)
		return
	}

	requestsTotal.WithLabelValues("success").Inc()
}

func main() {
	// Marca o servico como disponivel assim que ele sobe
	serviceUp.Set(1)

	http.HandleFunc("/projeto-korp", projetoKorpHandler)
	http.Handle("/metrics", promhttp.Handler())

	log.Println("http-server-projeto-korp iniciado na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
