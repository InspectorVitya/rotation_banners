package httpserver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type request struct {
	BannerID int `json:"banner_id"`
	SlotID   int `json:"slot_id"`
	GroupID  int `json:"group_id"`
}

func (s *Server) addBannerInSlot(w http.ResponseWriter, r *http.Request) {
	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		fmt.Println(err)
	}

	err := s.App.Storage.AddBanner(r.Context(), req.BannerID, req.SlotID)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(201)
	io.WriteString(w, "test")
}
