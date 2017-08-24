package alidns

import (
	api "github.com/denverdino/aliyungo/dns"
)

func (a *AlidnsProvider) getAllRecord() (resp *api.DescribeDomainRecordsResponse, err error) {
	resp, totalPage, err := a.getRecordPage(50, 1)
	if totalPage <= 1 {
		return resp, err
	} else {
		for i := 2; i <= totalPage; i++ {
			respAppend, _, err := a.getRecordPage(50, i)
			if err != nil {
				// skip
				continue
			} else {
				rec := resp.DomainRecords.Record
				recAppend := respAppend.DomainRecords.Record
				rec = append(rec, recAppend...)
				resp.DomainRecords.Record = rec
			}
		}
		return resp, err
	}
}

func (a *AlidnsProvider) getRecordPage(pageSize int, pageNumber int) (resp *api.DescribeDomainRecordsResponse, totalPage int, err error) {
	arg := api.DescribeDomainRecordsArgs{
		DomainName: a.rootDomainName,
	}
	arg.PageNumber = pageNumber
	arg.PageSize = pageSize
	resp, err = a.client.DescribeDomainRecords(&arg)
	if err != nil {
		return nil, 0, err
	}
	count := resp.TotalCount
	totalPage = count / pageSize
	if count%pageSize > 0 {
		totalPage++
	}
	return resp, totalPage, err
}
