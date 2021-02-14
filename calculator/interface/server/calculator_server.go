package server

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/snathale/challenge-hash/calculator/application/controller"
	"github.com/snathale/challenge-hash/calculator/interface/proto"
)

var CalculateDiscountError = errors.New("impossible calculate discount")

type CalculatorServer struct {
	controller controller.Controller
}

func NewCalculatorServer(controller controller.Controller) *CalculatorServer {
	return &CalculatorServer{
		controller: controller,
	}
}

func (c *CalculatorServer) GetProductDiscount(ctx context.Context, req *proto.Request) (*proto.Discount, error) {
	logrus.Info(fmt.Sprintf("calculating discount for user_id: %s and product_id: %s", req.UserId, req.ProductId))
	discount, err := c.controller.CalculateDiscount(req.UserId, req.ProductId)
	if err != nil {
		logrus.WithError(err).Warning(CalculateDiscountError)
		return nil, err
	}
	return &proto.Discount{
		Percentage:   discount.Percentage,
		ValueInCents: int32(discount.ValueInCents),
	}, nil
}
