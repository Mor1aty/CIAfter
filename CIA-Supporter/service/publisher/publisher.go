package publisher

import (
	"context"
	"github.com/micro/go-micro"
	"log"
	"moriaty.com/cia/cia-common/base/constant"
	pb "moriaty.com/cia/cia-common/proto/supporter/publisher"
	"moriaty.com/cia/cia-supporter/bean"
	"moriaty.com/cia/cia-supporter/dao/publisher"
	"strconv"
)

/**
 * @author 16计算机 Moriaty
 * @version 1.0
 * @copyright ：Moriaty 版权所有 © 2020
 * @date 2020/4/19 19:33
 * @Description TODO
 * 任务发布者服务 service
 */

type SupportPublisher struct{}

// 注册服务
func Register(service micro.Service) error {
	err := pb.RegisterSupporterPublisherHandler(service.Server(), new(SupportPublisher))
	if err != nil {
		log.Printf("register Publisher failed, err: %v", err)
		return err
	}
	return nil
}

// 获取所有业务
func (sp *SupportPublisher) FindAllBusiness(ctx context.Context, req *pb.FindAllBusinessReq, resp *pb.FindAllBusinessResp) (err error) {
	var tmpBusinesses []*bean.Business
	if req.All {
		tmpBusinesses, err = publisher.FindAllBusiness(2)
	} else if req.IsStop {
		tmpBusinesses, err = publisher.FindAllBusiness(1)
	} else {
		tmpBusinesses, err = publisher.FindAllBusiness(0)
	}
	if err != nil {
		log.Printf("find all business failed, err: %v", err)
		return err
	}
	if len(tmpBusinesses) == 0 {
		log.Printf("find all business not found")
		resp.Businesses = nil
		return nil
	}
	businesses := make([]*pb.Business, len(tmpBusinesses))
	var isStop bool
	for index, tmpBusiness := range tmpBusinesses {
		if tmpBusiness.IsStop == 1 {
			isStop = true
		}

		tmpBusinessTests, err := publisher.FindBusinessTestBySecret(tmpBusiness.Secret)
		businessTests := make([]*pb.BusinessTest, len(tmpBusinessTests))
		if err != nil {
			log.Printf("find business_test by secret [%s] failed, err: %v", tmpBusiness.Secret,
				err)
			businessTests = nil
		}
		if len(tmpBusinessTests) == 0 {
			log.Printf("find business_test by secret [%s] not found", tmpBusiness.Secret)
			businessTests = nil
		}
		for index2, tmpBusinessTest := range tmpBusinessTests {
			businessTests[index2] = &pb.BusinessTest{
				Id:         tmpBusinessTest.Id,
				Secret:     tmpBusinessTest.Secret,
				TestType:   tmpBusinessTest.TestType,
				TestOption: tmpBusinessTest.TestOption,
				UpdateTime: tmpBusinessTest.UpdateTime.Format(constant.DATE_TYPE_01),
			}
		}
		businesses[index] = &pb.Business{
			Secret:         tmpBusiness.Secret,
			SecretValue:    tmpBusiness.SecretValue,
			ApkName:        tmpBusiness.ApkName,
			PullFrequency:  tmpBusiness.PullFrequency,
			CurrentEdition: tmpBusiness.CurrentEdition,
			FileUrl:        tmpBusiness.FileUrl,
			IsStop:         isStop,
			UpdateTime:     tmpBusiness.UpdateTime.Format(constant.DATE_TYPE_01),
			BusinessTests:  businessTests,
		}
	}
	resp.Businesses = businesses
	log.Printf("find all business success")
	return nil
}

func (sp *SupportPublisher) UpdateBusinessEdition(ctx context.Context, req *pb.UpdateBusinessEditionReq, resp *pb.UpdateBusinessEditionResp) (err error) {
	edition := strconv.FormatFloat(float64(req.Edition), 'f', 2, 32)
	err = publisher.UpdateBusinessEdition(edition, req.Secret)
	if err != nil {
		log.Printf("update business edition failed, err: %v", err)
		return err
	}
	return nil
}
