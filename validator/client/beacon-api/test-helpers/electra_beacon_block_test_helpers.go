package test_helpers

import (
	"github.com/prysmaticlabs/prysm/v5/api/server/structs"
	"github.com/prysmaticlabs/prysm/v5/encoding/bytesutil"
	enginev1 "github.com/prysmaticlabs/prysm/v5/proto/engine/v1"
	ethpb "github.com/prysmaticlabs/prysm/v5/proto/prysm/v1alpha1"
)

func GenerateProtoElectraBeaconBlockContents() *ethpb.BeaconBlockContentsElectra {
	return &ethpb.BeaconBlockContentsElectra{
		Block: &ethpb.BeaconBlockElectra{
			Slot:          1,
			ProposerIndex: 2,
			ParentRoot:    FillByteSlice(32, 3),
			StateRoot:     FillByteSlice(32, 4),
			Body: &ethpb.BeaconBlockBodyElectra{
				RandaoReveal: FillByteSlice(96, 5),
				Eth1Data: &ethpb.Eth1Data{
					DepositRoot:  FillByteSlice(32, 6),
					DepositCount: 7,
					BlockHash:    FillByteSlice(32, 8),
				},
				Graffiti: FillByteSlice(32, 9),
				ProposerSlashings: []*ethpb.ProposerSlashing{
					{
						Header_1: &ethpb.SignedBeaconBlockHeader{
							Header: &ethpb.BeaconBlockHeader{
								Slot:          10,
								ProposerIndex: 11,
								ParentRoot:    FillByteSlice(32, 12),
								StateRoot:     FillByteSlice(32, 13),
								BodyRoot:      FillByteSlice(32, 14),
							},
							Signature: FillByteSlice(96, 15),
						},
						Header_2: &ethpb.SignedBeaconBlockHeader{
							Header: &ethpb.BeaconBlockHeader{
								Slot:          16,
								ProposerIndex: 17,
								ParentRoot:    FillByteSlice(32, 18),
								StateRoot:     FillByteSlice(32, 19),
								BodyRoot:      FillByteSlice(32, 20),
							},
							Signature: FillByteSlice(96, 21),
						},
					},
					{
						Header_1: &ethpb.SignedBeaconBlockHeader{
							Header: &ethpb.BeaconBlockHeader{
								Slot:          22,
								ProposerIndex: 23,
								ParentRoot:    FillByteSlice(32, 24),
								StateRoot:     FillByteSlice(32, 25),
								BodyRoot:      FillByteSlice(32, 26),
							},
							Signature: FillByteSlice(96, 27),
						},
						Header_2: &ethpb.SignedBeaconBlockHeader{
							Header: &ethpb.BeaconBlockHeader{
								Slot:          28,
								ProposerIndex: 29,
								ParentRoot:    FillByteSlice(32, 30),
								StateRoot:     FillByteSlice(32, 31),
								BodyRoot:      FillByteSlice(32, 32),
							},
							Signature: FillByteSlice(96, 33),
						},
					},
				},
				AttesterSlashings: []*ethpb.AttesterSlashingElectra{
					{
						Attestation_1: &ethpb.IndexedAttestationElectra{
							AttestingIndices: []uint64{34, 35},
							Data: &ethpb.AttestationData{
								Slot:            36,
								CommitteeIndex:  37,
								BeaconBlockRoot: FillByteSlice(32, 38),
								Source: &ethpb.Checkpoint{
									Epoch: 39,
									Root:  FillByteSlice(32, 40),
								},
								Target: &ethpb.Checkpoint{
									Epoch: 41,
									Root:  FillByteSlice(32, 42),
								},
							},
							Signature: FillByteSlice(96, 43),
						},
						Attestation_2: &ethpb.IndexedAttestationElectra{
							AttestingIndices: []uint64{44, 45},
							Data: &ethpb.AttestationData{
								Slot:            46,
								CommitteeIndex:  47,
								BeaconBlockRoot: FillByteSlice(32, 38),
								Source: &ethpb.Checkpoint{
									Epoch: 49,
									Root:  FillByteSlice(32, 50),
								},
								Target: &ethpb.Checkpoint{
									Epoch: 51,
									Root:  FillByteSlice(32, 52),
								},
							},
							Signature: FillByteSlice(96, 53),
						},
					},
				},
				Attestations: []*ethpb.AttestationElectra{
					{
						AggregationBits: FillByteSlice(4, 74),
						Data: &ethpb.AttestationData{
							Slot:            75,
							CommitteeIndex:  76,
							BeaconBlockRoot: FillByteSlice(32, 38),
							Source: &ethpb.Checkpoint{
								Epoch: 78,
								Root:  FillByteSlice(32, 79),
							},
							Target: &ethpb.Checkpoint{
								Epoch: 80,
								Root:  FillByteSlice(32, 81),
							},
						},
						Signature:     FillByteSlice(96, 82),
						CommitteeBits: FillByteSlice(8, 82),
					},
					{
						AggregationBits: FillByteSlice(4, 83),
						Data: &ethpb.AttestationData{
							Slot:            84,
							CommitteeIndex:  85,
							BeaconBlockRoot: FillByteSlice(32, 38),
							Source: &ethpb.Checkpoint{
								Epoch: 87,
								Root:  FillByteSlice(32, 88),
							},
							Target: &ethpb.Checkpoint{
								Epoch: 89,
								Root:  FillByteSlice(32, 90),
							},
						},
						Signature:     FillByteSlice(96, 91),
						CommitteeBits: FillByteSlice(8, 91),
					},
				},
				Deposits: []*ethpb.Deposit{
					{
						Proof: FillByteArraySlice(33, FillByteSlice(32, 92)),
						Data: &ethpb.Deposit_Data{
							PublicKey:             FillByteSlice(48, 94),
							WithdrawalCredentials: FillByteSlice(32, 95),
							Amount:                96,
							Signature:             FillByteSlice(96, 97),
						},
					},
					{
						Proof: FillByteArraySlice(33, FillByteSlice(32, 98)),
						Data: &ethpb.Deposit_Data{
							PublicKey:             FillByteSlice(48, 100),
							WithdrawalCredentials: FillByteSlice(32, 101),
							Amount:                102,
							Signature:             FillByteSlice(96, 103),
						},
					},
				},
				VoluntaryExits: []*ethpb.SignedVoluntaryExit{
					{
						Exit: &ethpb.VoluntaryExit{
							Epoch:          104,
							ValidatorIndex: 105,
						},
						Signature: FillByteSlice(96, 106),
					},
					{
						Exit: &ethpb.VoluntaryExit{
							Epoch:          107,
							ValidatorIndex: 108,
						},
						Signature: FillByteSlice(96, 109),
					},
				},
				SyncAggregate: &ethpb.SyncAggregate{
					SyncCommitteeBits:      FillByteSlice(64, 110),
					SyncCommitteeSignature: FillByteSlice(96, 111),
				},
				ExecutionPayload: &enginev1.ExecutionPayloadDeneb{
					ParentHash:    FillByteSlice(32, 112),
					FeeRecipient:  FillByteSlice(20, 113),
					StateRoot:     FillByteSlice(32, 114),
					ReceiptsRoot:  FillByteSlice(32, 115),
					LogsBloom:     FillByteSlice(256, 116),
					PrevRandao:    FillByteSlice(32, 117),
					BlockNumber:   118,
					GasLimit:      119,
					GasUsed:       120,
					Timestamp:     121,
					ExtraData:     FillByteSlice(32, 122),
					BaseFeePerGas: FillByteSlice(32, 123),
					BlockHash:     FillByteSlice(32, 124),
					Transactions: [][]byte{
						FillByteSlice(32, 125),
						FillByteSlice(32, 126),
					},
					Withdrawals: []*enginev1.Withdrawal{
						{
							Index:          127,
							ValidatorIndex: 128,
							Address:        FillByteSlice(20, 129),
							Amount:         130,
						},
						{
							Index:          131,
							ValidatorIndex: 132,
							Address:        FillByteSlice(20, 133),
							Amount:         134,
						},
					},
					BlobGasUsed:   135,
					ExcessBlobGas: 136,
				},
				BlsToExecutionChanges: []*ethpb.SignedBLSToExecutionChange{
					{
						Message: &ethpb.BLSToExecutionChange{
							ValidatorIndex:     137,
							FromBlsPubkey:      FillByteSlice(48, 138),
							ToExecutionAddress: FillByteSlice(20, 139),
						},
						Signature: FillByteSlice(96, 140),
					},
					{
						Message: &ethpb.BLSToExecutionChange{
							ValidatorIndex:     141,
							FromBlsPubkey:      FillByteSlice(48, 142),
							ToExecutionAddress: FillByteSlice(20, 143),
						},
						Signature: FillByteSlice(96, 144),
					},
				},
				BlobKzgCommitments: [][]byte{FillByteSlice(48, 145), FillByteSlice(48, 146)},
				ExecutionRequests: &enginev1.ExecutionRequests{
					Deposits: []*enginev1.DepositRequest{{
						Pubkey:                FillByteSlice(48, 147),
						WithdrawalCredentials: FillByteSlice(32, 148),
						Amount:                149,
						Signature:             FillByteSlice(96, 150),
						Index:                 151,
					}},
					Withdrawals: []*enginev1.WithdrawalRequest{
						{
							SourceAddress:   FillByteSlice(20, 152),
							ValidatorPubkey: FillByteSlice(48, 153),
							Amount:          154,
						},
					},
					Consolidations: []*enginev1.ConsolidationRequest{
						{
							SourceAddress: FillByteSlice(20, 155),
							SourcePubkey:  FillByteSlice(48, 156),
							TargetPubkey:  FillByteSlice(48, 157),
						},
					},
				},
			},
		},
		KzgProofs: [][]byte{FillByteSlice(48, 158)},
		Blobs:     [][]byte{FillByteSlice(131072, 159)},
	}
}

func GenerateProtoBlindedElectraBeaconBlock() *ethpb.BlindedBeaconBlockElectra {
	return &ethpb.BlindedBeaconBlockElectra{
		Slot:          1,
		ProposerIndex: 2,
		ParentRoot:    FillByteSlice(32, 3),
		StateRoot:     FillByteSlice(32, 4),
		Body: &ethpb.BlindedBeaconBlockBodyElectra{
			RandaoReveal: FillByteSlice(96, 5),
			Eth1Data: &ethpb.Eth1Data{
				DepositRoot:  FillByteSlice(32, 6),
				DepositCount: 7,
				BlockHash:    FillByteSlice(32, 8),
			},
			Graffiti: FillByteSlice(32, 9),
			ProposerSlashings: []*ethpb.ProposerSlashing{
				{
					Header_1: &ethpb.SignedBeaconBlockHeader{
						Header: &ethpb.BeaconBlockHeader{
							Slot:          10,
							ProposerIndex: 11,
							ParentRoot:    FillByteSlice(32, 12),
							StateRoot:     FillByteSlice(32, 13),
							BodyRoot:      FillByteSlice(32, 14),
						},
						Signature: FillByteSlice(96, 15),
					},
					Header_2: &ethpb.SignedBeaconBlockHeader{
						Header: &ethpb.BeaconBlockHeader{
							Slot:          16,
							ProposerIndex: 17,
							ParentRoot:    FillByteSlice(32, 18),
							StateRoot:     FillByteSlice(32, 19),
							BodyRoot:      FillByteSlice(32, 20),
						},
						Signature: FillByteSlice(96, 21),
					},
				},
				{
					Header_1: &ethpb.SignedBeaconBlockHeader{
						Header: &ethpb.BeaconBlockHeader{
							Slot:          22,
							ProposerIndex: 23,
							ParentRoot:    FillByteSlice(32, 24),
							StateRoot:     FillByteSlice(32, 25),
							BodyRoot:      FillByteSlice(32, 26),
						},
						Signature: FillByteSlice(96, 27),
					},
					Header_2: &ethpb.SignedBeaconBlockHeader{
						Header: &ethpb.BeaconBlockHeader{
							Slot:          28,
							ProposerIndex: 29,
							ParentRoot:    FillByteSlice(32, 30),
							StateRoot:     FillByteSlice(32, 31),
							BodyRoot:      FillByteSlice(32, 32),
						},
						Signature: FillByteSlice(96, 33),
					},
				},
			},
			AttesterSlashings: []*ethpb.AttesterSlashingElectra{
				{
					Attestation_1: &ethpb.IndexedAttestationElectra{
						AttestingIndices: []uint64{34, 35},
						Data: &ethpb.AttestationData{
							Slot:            36,
							CommitteeIndex:  37,
							BeaconBlockRoot: FillByteSlice(32, 38),
							Source: &ethpb.Checkpoint{
								Epoch: 39,
								Root:  FillByteSlice(32, 40),
							},
							Target: &ethpb.Checkpoint{
								Epoch: 41,
								Root:  FillByteSlice(32, 42),
							},
						},
						Signature: FillByteSlice(96, 43),
					},
					Attestation_2: &ethpb.IndexedAttestationElectra{
						AttestingIndices: []uint64{44, 45},
						Data: &ethpb.AttestationData{
							Slot:            46,
							CommitteeIndex:  47,
							BeaconBlockRoot: FillByteSlice(32, 38),
							Source: &ethpb.Checkpoint{
								Epoch: 49,
								Root:  FillByteSlice(32, 50),
							},
							Target: &ethpb.Checkpoint{
								Epoch: 51,
								Root:  FillByteSlice(32, 52),
							},
						},
						Signature: FillByteSlice(96, 53),
					},
				},
			},
			Attestations: []*ethpb.AttestationElectra{
				{
					AggregationBits: FillByteSlice(4, 74),
					Data: &ethpb.AttestationData{
						Slot:            75,
						CommitteeIndex:  76,
						BeaconBlockRoot: FillByteSlice(32, 38),
						Source: &ethpb.Checkpoint{
							Epoch: 78,
							Root:  FillByteSlice(32, 79),
						},
						Target: &ethpb.Checkpoint{
							Epoch: 80,
							Root:  FillByteSlice(32, 81),
						},
					},
					Signature:     FillByteSlice(96, 82),
					CommitteeBits: FillByteSlice(8, 82),
				},
				{
					AggregationBits: FillByteSlice(4, 83),
					Data: &ethpb.AttestationData{
						Slot:            84,
						CommitteeIndex:  85,
						BeaconBlockRoot: FillByteSlice(32, 38),
						Source: &ethpb.Checkpoint{
							Epoch: 87,
							Root:  FillByteSlice(32, 88),
						},
						Target: &ethpb.Checkpoint{
							Epoch: 89,
							Root:  FillByteSlice(32, 90),
						},
					},
					Signature:     FillByteSlice(96, 91),
					CommitteeBits: FillByteSlice(8, 91),
				},
			},
			Deposits: []*ethpb.Deposit{
				{
					Proof: FillByteArraySlice(33, FillByteSlice(32, 92)),
					Data: &ethpb.Deposit_Data{
						PublicKey:             FillByteSlice(48, 94),
						WithdrawalCredentials: FillByteSlice(32, 95),
						Amount:                96,
						Signature:             FillByteSlice(96, 97),
					},
				},
				{
					Proof: FillByteArraySlice(33, FillByteSlice(32, 98)),
					Data: &ethpb.Deposit_Data{
						PublicKey:             FillByteSlice(48, 100),
						WithdrawalCredentials: FillByteSlice(32, 101),
						Amount:                102,
						Signature:             FillByteSlice(96, 103),
					},
				},
			},
			VoluntaryExits: []*ethpb.SignedVoluntaryExit{
				{
					Exit: &ethpb.VoluntaryExit{
						Epoch:          104,
						ValidatorIndex: 105,
					},
					Signature: FillByteSlice(96, 106),
				},
				{
					Exit: &ethpb.VoluntaryExit{
						Epoch:          107,
						ValidatorIndex: 108,
					},
					Signature: FillByteSlice(96, 109),
				},
			},
			SyncAggregate: &ethpb.SyncAggregate{
				SyncCommitteeBits:      FillByteSlice(64, 110),
				SyncCommitteeSignature: FillByteSlice(96, 111),
			},
			ExecutionPayloadHeader: &enginev1.ExecutionPayloadHeaderDeneb{
				ParentHash:       FillByteSlice(32, 112),
				FeeRecipient:     FillByteSlice(20, 113),
				StateRoot:        FillByteSlice(32, 114),
				ReceiptsRoot:     FillByteSlice(32, 115),
				LogsBloom:        FillByteSlice(256, 116),
				PrevRandao:       FillByteSlice(32, 117),
				BlockNumber:      118,
				GasLimit:         119,
				GasUsed:          120,
				Timestamp:        121,
				ExtraData:        FillByteSlice(32, 122),
				BaseFeePerGas:    FillByteSlice(32, 123),
				BlockHash:        FillByteSlice(32, 124),
				TransactionsRoot: FillByteSlice(32, 125),
				WithdrawalsRoot:  FillByteSlice(32, 126),
				BlobGasUsed:      127,
				ExcessBlobGas:    128,
			},
			BlsToExecutionChanges: []*ethpb.SignedBLSToExecutionChange{
				{
					Message: &ethpb.BLSToExecutionChange{
						ValidatorIndex:     129,
						FromBlsPubkey:      FillByteSlice(48, 130),
						ToExecutionAddress: FillByteSlice(20, 131),
					},
					Signature: FillByteSlice(96, 132),
				},
				{
					Message: &ethpb.BLSToExecutionChange{
						ValidatorIndex:     133,
						FromBlsPubkey:      FillByteSlice(48, 134),
						ToExecutionAddress: FillByteSlice(20, 135),
					},
					Signature: FillByteSlice(96, 136),
				},
			},
			BlobKzgCommitments: [][]byte{FillByteSlice(48, 137), FillByteSlice(48, 138)},
			ExecutionRequests: &enginev1.ExecutionRequests{
				Deposits: []*enginev1.DepositRequest{{
					Pubkey:                FillByteSlice(48, 139),
					WithdrawalCredentials: FillByteSlice(32, 140),
					Amount:                141,
					Signature:             FillByteSlice(96, 142),
					Index:                 143,
				}},
				Withdrawals: []*enginev1.WithdrawalRequest{
					{
						SourceAddress:   FillByteSlice(20, 144),
						ValidatorPubkey: FillByteSlice(48, 145),
						Amount:          146,
					},
				},
				Consolidations: []*enginev1.ConsolidationRequest{
					{
						SourceAddress: FillByteSlice(20, 147),
						SourcePubkey:  FillByteSlice(48, 148),
						TargetPubkey:  FillByteSlice(48, 149),
					},
				},
			},
		},
	}
}

func GenerateJsonElectraBeaconBlockContents() *structs.BeaconBlockContentsElectra {
	return &structs.BeaconBlockContentsElectra{
		Block: &structs.BeaconBlockElectra{
			Slot:          "1",
			ProposerIndex: "2",
			ParentRoot:    FillEncodedByteSlice(32, 3),
			StateRoot:     FillEncodedByteSlice(32, 4),
			Body: &structs.BeaconBlockBodyElectra{
				RandaoReveal: FillEncodedByteSlice(96, 5),
				Eth1Data: &structs.Eth1Data{
					DepositRoot:  FillEncodedByteSlice(32, 6),
					DepositCount: "7",
					BlockHash:    FillEncodedByteSlice(32, 8),
				},
				Graffiti: FillEncodedByteSlice(32, 9),
				ProposerSlashings: []*structs.ProposerSlashing{
					{
						SignedHeader1: &structs.SignedBeaconBlockHeader{
							Message: &structs.BeaconBlockHeader{
								Slot:          "10",
								ProposerIndex: "11",
								ParentRoot:    FillEncodedByteSlice(32, 12),
								StateRoot:     FillEncodedByteSlice(32, 13),
								BodyRoot:      FillEncodedByteSlice(32, 14),
							},
							Signature: FillEncodedByteSlice(96, 15),
						},
						SignedHeader2: &structs.SignedBeaconBlockHeader{
							Message: &structs.BeaconBlockHeader{
								Slot:          "16",
								ProposerIndex: "17",
								ParentRoot:    FillEncodedByteSlice(32, 18),
								StateRoot:     FillEncodedByteSlice(32, 19),
								BodyRoot:      FillEncodedByteSlice(32, 20),
							},
							Signature: FillEncodedByteSlice(96, 21),
						},
					},
					{
						SignedHeader1: &structs.SignedBeaconBlockHeader{
							Message: &structs.BeaconBlockHeader{
								Slot:          "22",
								ProposerIndex: "23",
								ParentRoot:    FillEncodedByteSlice(32, 24),
								StateRoot:     FillEncodedByteSlice(32, 25),
								BodyRoot:      FillEncodedByteSlice(32, 26),
							},
							Signature: FillEncodedByteSlice(96, 27),
						},
						SignedHeader2: &structs.SignedBeaconBlockHeader{
							Message: &structs.BeaconBlockHeader{
								Slot:          "28",
								ProposerIndex: "29",
								ParentRoot:    FillEncodedByteSlice(32, 30),
								StateRoot:     FillEncodedByteSlice(32, 31),
								BodyRoot:      FillEncodedByteSlice(32, 32),
							},
							Signature: FillEncodedByteSlice(96, 33),
						},
					},
				},
				AttesterSlashings: []*structs.AttesterSlashingElectra{
					{
						Attestation1: &structs.IndexedAttestationElectra{
							AttestingIndices: []string{"34", "35"},
							Data: &structs.AttestationData{
								Slot:            "36",
								CommitteeIndex:  "37",
								BeaconBlockRoot: FillEncodedByteSlice(32, 38),
								Source: &structs.Checkpoint{
									Epoch: "39",
									Root:  FillEncodedByteSlice(32, 40),
								},
								Target: &structs.Checkpoint{
									Epoch: "41",
									Root:  FillEncodedByteSlice(32, 42),
								},
							},
							Signature: FillEncodedByteSlice(96, 43),
						},
						Attestation2: &structs.IndexedAttestationElectra{
							AttestingIndices: []string{"44", "45"},
							Data: &structs.AttestationData{
								Slot:            "46",
								CommitteeIndex:  "47",
								BeaconBlockRoot: FillEncodedByteSlice(32, 38),
								Source: &structs.Checkpoint{
									Epoch: "49",
									Root:  FillEncodedByteSlice(32, 50),
								},
								Target: &structs.Checkpoint{
									Epoch: "51",
									Root:  FillEncodedByteSlice(32, 52),
								},
							},
							Signature: FillEncodedByteSlice(96, 53),
						},
					},
				},
				Attestations: []*structs.AttestationElectra{
					{
						AggregationBits: FillEncodedByteSlice(4, 74),
						Data: &structs.AttestationData{
							Slot:            "75",
							CommitteeIndex:  "76",
							BeaconBlockRoot: FillEncodedByteSlice(32, 38),
							Source: &structs.Checkpoint{
								Epoch: "78",
								Root:  FillEncodedByteSlice(32, 79),
							},
							Target: &structs.Checkpoint{
								Epoch: "80",
								Root:  FillEncodedByteSlice(32, 81),
							},
						},
						Signature:     FillEncodedByteSlice(96, 82),
						CommitteeBits: FillEncodedByteSlice(8, 82),
					},
					{
						AggregationBits: FillEncodedByteSlice(4, 83),
						Data: &structs.AttestationData{
							Slot:            "84",
							CommitteeIndex:  "85",
							BeaconBlockRoot: FillEncodedByteSlice(32, 38),
							Source: &structs.Checkpoint{
								Epoch: "87",
								Root:  FillEncodedByteSlice(32, 88),
							},
							Target: &structs.Checkpoint{
								Epoch: "89",
								Root:  FillEncodedByteSlice(32, 90),
							},
						},
						Signature:     FillEncodedByteSlice(96, 91),
						CommitteeBits: FillEncodedByteSlice(8, 91),
					},
				},
				Deposits: []*structs.Deposit{
					{
						Proof: FillEncodedByteArraySlice(33, FillEncodedByteSlice(32, 92)),
						Data: &structs.DepositData{
							Pubkey:                FillEncodedByteSlice(48, 94),
							WithdrawalCredentials: FillEncodedByteSlice(32, 95),
							Amount:                "96",
							Signature:             FillEncodedByteSlice(96, 97),
						},
					},
					{
						Proof: FillEncodedByteArraySlice(33, FillEncodedByteSlice(32, 98)),
						Data: &structs.DepositData{
							Pubkey:                FillEncodedByteSlice(48, 100),
							WithdrawalCredentials: FillEncodedByteSlice(32, 101),
							Amount:                "102",
							Signature:             FillEncodedByteSlice(96, 103),
						},
					},
				},
				VoluntaryExits: []*structs.SignedVoluntaryExit{
					{
						Message: &structs.VoluntaryExit{
							Epoch:          "104",
							ValidatorIndex: "105",
						},
						Signature: FillEncodedByteSlice(96, 106),
					},
					{
						Message: &structs.VoluntaryExit{
							Epoch:          "107",
							ValidatorIndex: "108",
						},
						Signature: FillEncodedByteSlice(96, 109),
					},
				},
				SyncAggregate: &structs.SyncAggregate{
					SyncCommitteeBits:      FillEncodedByteSlice(64, 110),
					SyncCommitteeSignature: FillEncodedByteSlice(96, 111),
				},
				ExecutionPayload: &structs.ExecutionPayloadDeneb{
					ParentHash:    FillEncodedByteSlice(32, 112),
					FeeRecipient:  FillEncodedByteSlice(20, 113),
					StateRoot:     FillEncodedByteSlice(32, 114),
					ReceiptsRoot:  FillEncodedByteSlice(32, 115),
					LogsBloom:     FillEncodedByteSlice(256, 116),
					PrevRandao:    FillEncodedByteSlice(32, 117),
					BlockNumber:   "118",
					GasLimit:      "119",
					GasUsed:       "120",
					Timestamp:     "121",
					ExtraData:     FillEncodedByteSlice(32, 122),
					BaseFeePerGas: bytesutil.LittleEndianBytesToBigInt(FillByteSlice(32, 123)).String(),
					BlockHash:     FillEncodedByteSlice(32, 124),
					Transactions: []string{
						FillEncodedByteSlice(32, 125),
						FillEncodedByteSlice(32, 126),
					},
					Withdrawals: []*structs.Withdrawal{
						{
							WithdrawalIndex:  "127",
							ValidatorIndex:   "128",
							ExecutionAddress: FillEncodedByteSlice(20, 129),
							Amount:           "130",
						},
						{
							WithdrawalIndex:  "131",
							ValidatorIndex:   "132",
							ExecutionAddress: FillEncodedByteSlice(20, 133),
							Amount:           "134",
						},
					},
					BlobGasUsed:   "135",
					ExcessBlobGas: "136",
				},
				BLSToExecutionChanges: []*structs.SignedBLSToExecutionChange{
					{
						Message: &structs.BLSToExecutionChange{
							ValidatorIndex:     "137",
							FromBLSPubkey:      FillEncodedByteSlice(48, 138),
							ToExecutionAddress: FillEncodedByteSlice(20, 139),
						},
						Signature: FillEncodedByteSlice(96, 140),
					},
					{
						Message: &structs.BLSToExecutionChange{
							ValidatorIndex:     "141",
							FromBLSPubkey:      FillEncodedByteSlice(48, 142),
							ToExecutionAddress: FillEncodedByteSlice(20, 143),
						},
						Signature: FillEncodedByteSlice(96, 144),
					},
				},
				BlobKzgCommitments: []string{FillEncodedByteSlice(48, 145), FillEncodedByteSlice(48, 146)},
				ExecutionRequests: &structs.ExecutionRequests{
					Deposits: []*structs.DepositRequest{{
						Pubkey:                FillEncodedByteSlice(48, 147),
						WithdrawalCredentials: FillEncodedByteSlice(32, 148),
						Amount:                "149",
						Signature:             FillEncodedByteSlice(96, 150),
						Index:                 "151",
					}},
					Withdrawals: []*structs.WithdrawalRequest{
						{
							SourceAddress:   FillEncodedByteSlice(20, 152),
							ValidatorPubkey: FillEncodedByteSlice(48, 153),
							Amount:          "154",
						},
					},
					Consolidations: []*structs.ConsolidationRequest{
						{
							SourceAddress: FillEncodedByteSlice(20, 155),
							SourcePubkey:  FillEncodedByteSlice(48, 156),
							TargetPubkey:  FillEncodedByteSlice(48, 157),
						},
					},
				},
			},
		},
		KzgProofs: []string{FillEncodedByteSlice(48, 158)},
		Blobs:     []string{FillEncodedByteSlice(131072, 159)},
	}
}

func GenerateJsonBlindedElectraBeaconBlock() *structs.BlindedBeaconBlockElectra {
	return &structs.BlindedBeaconBlockElectra{
		Slot:          "1",
		ProposerIndex: "2",
		ParentRoot:    FillEncodedByteSlice(32, 3),
		StateRoot:     FillEncodedByteSlice(32, 4),
		Body: &structs.BlindedBeaconBlockBodyElectra{
			RandaoReveal: FillEncodedByteSlice(96, 5),
			Eth1Data: &structs.Eth1Data{
				DepositRoot:  FillEncodedByteSlice(32, 6),
				DepositCount: "7",
				BlockHash:    FillEncodedByteSlice(32, 8),
			},
			Graffiti: FillEncodedByteSlice(32, 9),
			ProposerSlashings: []*structs.ProposerSlashing{
				{
					SignedHeader1: &structs.SignedBeaconBlockHeader{
						Message: &structs.BeaconBlockHeader{
							Slot:          "10",
							ProposerIndex: "11",
							ParentRoot:    FillEncodedByteSlice(32, 12),
							StateRoot:     FillEncodedByteSlice(32, 13),
							BodyRoot:      FillEncodedByteSlice(32, 14),
						},
						Signature: FillEncodedByteSlice(96, 15),
					},
					SignedHeader2: &structs.SignedBeaconBlockHeader{
						Message: &structs.BeaconBlockHeader{
							Slot:          "16",
							ProposerIndex: "17",
							ParentRoot:    FillEncodedByteSlice(32, 18),
							StateRoot:     FillEncodedByteSlice(32, 19),
							BodyRoot:      FillEncodedByteSlice(32, 20),
						},
						Signature: FillEncodedByteSlice(96, 21),
					},
				},
				{
					SignedHeader1: &structs.SignedBeaconBlockHeader{
						Message: &structs.BeaconBlockHeader{
							Slot:          "22",
							ProposerIndex: "23",
							ParentRoot:    FillEncodedByteSlice(32, 24),
							StateRoot:     FillEncodedByteSlice(32, 25),
							BodyRoot:      FillEncodedByteSlice(32, 26),
						},
						Signature: FillEncodedByteSlice(96, 27),
					},
					SignedHeader2: &structs.SignedBeaconBlockHeader{
						Message: &structs.BeaconBlockHeader{
							Slot:          "28",
							ProposerIndex: "29",
							ParentRoot:    FillEncodedByteSlice(32, 30),
							StateRoot:     FillEncodedByteSlice(32, 31),
							BodyRoot:      FillEncodedByteSlice(32, 32),
						},
						Signature: FillEncodedByteSlice(96, 33),
					},
				},
			},
			AttesterSlashings: []*structs.AttesterSlashingElectra{
				{
					Attestation1: &structs.IndexedAttestationElectra{
						AttestingIndices: []string{"34", "35"},
						Data: &structs.AttestationData{
							Slot:            "36",
							CommitteeIndex:  "37",
							BeaconBlockRoot: FillEncodedByteSlice(32, 38),
							Source: &structs.Checkpoint{
								Epoch: "39",
								Root:  FillEncodedByteSlice(32, 40),
							},
							Target: &structs.Checkpoint{
								Epoch: "41",
								Root:  FillEncodedByteSlice(32, 42),
							},
						},
						Signature: FillEncodedByteSlice(96, 43),
					},
					Attestation2: &structs.IndexedAttestationElectra{
						AttestingIndices: []string{"44", "45"},
						Data: &structs.AttestationData{
							Slot:            "46",
							CommitteeIndex:  "47",
							BeaconBlockRoot: FillEncodedByteSlice(32, 38),
							Source: &structs.Checkpoint{
								Epoch: "49",
								Root:  FillEncodedByteSlice(32, 50),
							},
							Target: &structs.Checkpoint{
								Epoch: "51",
								Root:  FillEncodedByteSlice(32, 52),
							},
						},
						Signature: FillEncodedByteSlice(96, 53),
					},
				},
			},
			Attestations: []*structs.AttestationElectra{
				{
					AggregationBits: FillEncodedByteSlice(4, 74),
					Data: &structs.AttestationData{
						Slot:            "75",
						CommitteeIndex:  "76",
						BeaconBlockRoot: FillEncodedByteSlice(32, 38),
						Source: &structs.Checkpoint{
							Epoch: "78",
							Root:  FillEncodedByteSlice(32, 79),
						},
						Target: &structs.Checkpoint{
							Epoch: "80",
							Root:  FillEncodedByteSlice(32, 81),
						},
					},
					Signature:     FillEncodedByteSlice(96, 82),
					CommitteeBits: FillEncodedByteSlice(8, 82),
				},
				{
					AggregationBits: FillEncodedByteSlice(4, 83),
					Data: &structs.AttestationData{
						Slot:            "84",
						CommitteeIndex:  "85",
						BeaconBlockRoot: FillEncodedByteSlice(32, 38),
						Source: &structs.Checkpoint{
							Epoch: "87",
							Root:  FillEncodedByteSlice(32, 88),
						},
						Target: &structs.Checkpoint{
							Epoch: "89",
							Root:  FillEncodedByteSlice(32, 90),
						},
					},
					Signature:     FillEncodedByteSlice(96, 91),
					CommitteeBits: FillEncodedByteSlice(8, 91),
				},
			},
			Deposits: []*structs.Deposit{
				{
					Proof: FillEncodedByteArraySlice(33, FillEncodedByteSlice(32, 92)),
					Data: &structs.DepositData{
						Pubkey:                FillEncodedByteSlice(48, 94),
						WithdrawalCredentials: FillEncodedByteSlice(32, 95),
						Amount:                "96",
						Signature:             FillEncodedByteSlice(96, 97),
					},
				},
				{
					Proof: FillEncodedByteArraySlice(33, FillEncodedByteSlice(32, 98)),
					Data: &structs.DepositData{
						Pubkey:                FillEncodedByteSlice(48, 100),
						WithdrawalCredentials: FillEncodedByteSlice(32, 101),
						Amount:                "102",
						Signature:             FillEncodedByteSlice(96, 103),
					},
				},
			},
			VoluntaryExits: []*structs.SignedVoluntaryExit{
				{
					Message: &structs.VoluntaryExit{
						Epoch:          "104",
						ValidatorIndex: "105",
					},
					Signature: FillEncodedByteSlice(96, 106),
				},
				{
					Message: &structs.VoluntaryExit{
						Epoch:          "107",
						ValidatorIndex: "108",
					},
					Signature: FillEncodedByteSlice(96, 109),
				},
			},
			SyncAggregate: &structs.SyncAggregate{
				SyncCommitteeBits:      FillEncodedByteSlice(64, 110),
				SyncCommitteeSignature: FillEncodedByteSlice(96, 111),
			},
			ExecutionPayloadHeader: &structs.ExecutionPayloadHeaderDeneb{
				ParentHash:       FillEncodedByteSlice(32, 112),
				FeeRecipient:     FillEncodedByteSlice(20, 113),
				StateRoot:        FillEncodedByteSlice(32, 114),
				ReceiptsRoot:     FillEncodedByteSlice(32, 115),
				LogsBloom:        FillEncodedByteSlice(256, 116),
				PrevRandao:       FillEncodedByteSlice(32, 117),
				BlockNumber:      "118",
				GasLimit:         "119",
				GasUsed:          "120",
				Timestamp:        "121",
				ExtraData:        FillEncodedByteSlice(32, 122),
				BaseFeePerGas:    bytesutil.LittleEndianBytesToBigInt(FillByteSlice(32, 123)).String(),
				BlockHash:        FillEncodedByteSlice(32, 124),
				TransactionsRoot: FillEncodedByteSlice(32, 125),
				WithdrawalsRoot:  FillEncodedByteSlice(32, 126),
				BlobGasUsed:      "127",
				ExcessBlobGas:    "128",
			},
			BLSToExecutionChanges: []*structs.SignedBLSToExecutionChange{
				{
					Message: &structs.BLSToExecutionChange{
						ValidatorIndex:     "129",
						FromBLSPubkey:      FillEncodedByteSlice(48, 130),
						ToExecutionAddress: FillEncodedByteSlice(20, 131),
					},
					Signature: FillEncodedByteSlice(96, 132),
				},
				{
					Message: &structs.BLSToExecutionChange{
						ValidatorIndex:     "133",
						FromBLSPubkey:      FillEncodedByteSlice(48, 134),
						ToExecutionAddress: FillEncodedByteSlice(20, 135),
					},
					Signature: FillEncodedByteSlice(96, 136),
				},
			},
			BlobKzgCommitments: []string{FillEncodedByteSlice(48, 137), FillEncodedByteSlice(48, 138)},
			ExecutionRequests: &structs.ExecutionRequests{
				Deposits: []*structs.DepositRequest{{
					Pubkey:                FillEncodedByteSlice(48, 139),
					WithdrawalCredentials: FillEncodedByteSlice(32, 140),
					Amount:                "141",
					Signature:             FillEncodedByteSlice(96, 142),
					Index:                 "143",
				}},
				Withdrawals: []*structs.WithdrawalRequest{
					{
						SourceAddress:   FillEncodedByteSlice(20, 144),
						ValidatorPubkey: FillEncodedByteSlice(48, 145),
						Amount:          "146",
					},
				},
				Consolidations: []*structs.ConsolidationRequest{
					{
						SourceAddress: FillEncodedByteSlice(20, 147),
						SourcePubkey:  FillEncodedByteSlice(48, 148),
						TargetPubkey:  FillEncodedByteSlice(48, 149),
					},
				},
			},
		},
	}
}
