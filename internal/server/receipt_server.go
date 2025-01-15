package server

import (
	"net/http"
	"receipt_processor/internal/receipt"
	"receipt_processor/internal/server/models"

	"github.com/labstack/echo/v4"
)

type ReceiptServer struct {
	receiptSvc *receipt.ReceiptService
}

func (s *ReceiptServer) GetReceiptsIdPoints(ctx echo.Context, id string) error {

	points, err := s.receiptSvc.LookupPoints(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"description": "No receipt found for that ID."})
	}

	return ctx.JSON(http.StatusOK, models.CreatePointsResponse(points))
}

func (s *ReceiptServer) PostReceiptsProcess(ctx echo.Context) error {

	data := &models.Receipt{}

	if err := ctx.Bind(data); err != nil {
		// return err
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "The receipt is invalid."})
	}

	id, err := s.receiptSvc.ProcessReceipt(data)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "The receipt is invalid."})
	}
	return ctx.JSON(http.StatusOK, map[string]string{"id": id})
}

func New(receiptSvc *receipt.ReceiptService) (*ReceiptServer, error) {

	server := &ReceiptServer{
		receiptSvc: receiptSvc,
	}

	return server, nil
}
