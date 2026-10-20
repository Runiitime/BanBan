package handlers

import (
	"backend/internal/pgs/queries"
	"backend/models"
	"backend/tools/msg"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type BoardHandler struct {
	queries *queries.BoardQuery
}

func NewBoardHandler(q *queries.BoardQuery) *BoardHandler {
	return &BoardHandler{queries: q}
}

// HandleCreateBoard Create board
//
//	@Summary      Create a board
//	@Description  create a board
//	@Tags		  board
//	@Accept       json
//	@Produce      json
//	@Param		  board	body	 models.Board	true	"board info"
//	@Success      201  {object}  models.Board
//	@Failure      500  {object}  models.ErrorDTO
//	@Router       /boards [post]
func (h *BoardHandler) HandleCreateBoard(w http.ResponseWriter, r *http.Request) {
	var data models.Board

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&data); err != nil {
		msg.DoError(err, w, http.StatusInternalServerError)
		return
	}

	log.Println("r.Body", r.Body)
	id, err := h.queries.Create(r.Context(), data.Title)
	if err != nil {
		msg.DoError(err, w, http.StatusInternalServerError)
		return
	}

	log.Println(msg.BoardWasCreated)

	b, err := json.MarshalIndent(id, "", "    ")

	if err != nil {
		msg.DoError(err, w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		msg.DoError(err, w, http.StatusInternalServerError)
		return
	}
}

// HandleDeleteBoard Delete board
//
//	@Summary      Delete a board
//	@Description  delete a board by id
//	@Tags		  board
//	@Param		  id	path	 int	true	"board ID"	Format(int)
//	@Success      204  {}  No content
//	@Failure      404  {object}  models.ErrorDTO
//	@Failure      500  {object}  models.ErrorDTO
//	@Router       /boards/{id}    [delete]
func (h *BoardHandler) HandleDeleteBoard(w http.ResponseWriter, r *http.Request) {
	id, _ := mux.Vars(r)["id"]
	boardID, err := strconv.Atoi(id)
	if err != nil {
		msg.DoError(err, w, http.StatusNotFound)
		return
	}

	if err := h.queries.DeleteBoard(r.Context(), boardID); err != nil {
		if errors.Is(err, errors.New("no board found")) {
			msg.DoError(err, w, http.StatusNotFound)
			return
		}
		msg.DoError(err, w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// HandleGetBoard Get board by id
//
//		@Summary      Get a board by id
//		@Description  get a board by id
//		@Tags		  board
//		@Produce      json
//	 	@Param 		  id path int true "board ID"	Format(int)
//		@Success      200  {object}  models.Board
//		@Failure      404  {object}  models.ErrorDTO
//		@Failure      500  {object}  models.ErrorDTO
//		@Router       /boards/{id} [get]
func (h *BoardHandler) HandleGetBoard(w http.ResponseWriter, r *http.Request) {
	id, _ := mux.Vars(r)["id"]

	n, _ := strconv.Atoi(id)
	board, err := h.queries.GetBoard(r.Context(), n)

	if err != nil {
		msg.DoError(err, w)
		return
	}

	b, err := json.MarshalIndent(board, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}

	if _, err := w.Write(b); err != nil {
		msg.DoError(err, w, http.StatusInternalServerError)
		return
	}
}

// HandleUpdateBoard Update board
//
//	@Summary      Change a book status "completed"
//	@Description  change a book status "completed"
//	@Tags		  book
//	@Accept       json
//	@Param 		  id path int true "Book ID"	Format(int)
//	@Param 		  completed body  models.BoardStatusDTO true "completed value"
//	@Success      204  {}       No content
//	@Failure      404  {object}  models.ErrorDTO
//	@Failure      500  {object}  models.ErrorDTO
//	@Router       /books/{id} [patch]
func (h *BoardHandler) HandleUpdateBoard(w http.ResponseWriter, r *http.Request) {
	//id := mux.Vars(r)["id"]
	//
	//decoder := json.NewDecoder(r.Body)
	//
	//var status models.BoardStatusDTO
	//
	//if err := decoder.Decode(&status); err != nil {
	//	msg.DoError(err, w, http.StatusInternalServerError)
	//	return
	//}
	//
	//n, err := strconv.Atoi(id)
	//
	//if err != nil {
	//	msg.DoError(err, w)
	//	return
	//}
	//
	//if err := h.queries.UpdateBookStatus(r.Context(), n, status.Completed); err != nil {
	//	if errors.Is(err, bookMSG.ErrBookNotFound) {
	//		msg.DoError(bookMSG.ErrBookNotFound, w, http.StatusNotFound)
	//		return
	//	}
	//	msg.DoError(err, w, http.StatusInternalServerError)
	//	return
	//}
	//
	//w.WriteHeader(http.StatusNoContent)
}

// HandleGetAllBoards Get all boards
//
//	@Summary      Get all boards
//	@Description  get all boards
//	@Tags		  boards
//	@Produce      json
//	@Success      200  {array}   []models.Board
//	@Failure      500  {object}  models.ErrorDTO
//	@Router       /boards [get]
func (h *BoardHandler) HandleGetAllBoards(w http.ResponseWriter, r *http.Request) {
	boards, err := h.queries.GetBoards(r.Context())

	if err != nil {
		msg.DoError(err, w, http.StatusInternalServerError)
		return
	}

	b, err := json.MarshalIndent(boards, "", "    ")

	if err != nil {
		msg.DoError(err, w, http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		msg.DoError(err, w, http.StatusInternalServerError)
		return
	}
}
