/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package tcpserver

import (
	"encoding/base64"
	"errors"
	"time"

	"github.com/IBAX-io/go-ibax/packages/conf/syspar"
	"github.com/IBAX-io/go-ibax/packages/consts"
	"github.com/IBAX-io/go-ibax/packages/crypto"
	"github.com/IBAX-io/go-ibax/packages/crypto/ecies"
	"github.com/IBAX-io/go-ibax/packages/model"
	"github.com/IBAX-io/go-ibax/packages/network"
	"github.com/IBAX-io/go-ibax/packages/utils"

	log "github.com/sirupsen/logrus"
)

func Type88(r *network.PrivateDateRequest) (*network.PrivateDateResponse, error) {
	node_pri := syspar.GetNodePrivKey()
	data, err := ecies.EccDeCrypto(r.Data, node_pri)
	if err != nil {
		log.WithError(err)
		return nil, err
	}
	//hash, err := crypto.HashHex(r.Data)
	hash, err := crypto.HashHex(data)
	if err != nil {
		log.WithError(err)
		return nil, err
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("EccCryptoKey error")
		return nil, err
	}
	encodeDataString := base64.StdEncoding.EncodeToString(eccData)
	////

	privatePackets := model.PrivatePackets{
		Hash: hash,
		//Data: r.Data,
		//
		//Data: data,
		Data: []byte(encodeDataString),
		Time: time.Now().Unix(),
	}

	err = privatePackets.Create()
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Create PrivatePackets table record error")
		return nil, err
	}

	return resp, nil
}
