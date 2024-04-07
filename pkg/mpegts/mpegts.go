// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package mpegts

// MPEG: Moving Picture Experts Group

// FixedFragmentHeader 每个TS文件都以固定的PAT，PMT开始
// TODO(chef): [refactor] 考虑去掉这个fixed变量，目前只在测试中使用 [202303]
var FixedFragmentHeader = []byte{
	/* TS */
	0x47, 0x40, 0x00, 0x10, 0x00,
	/* PSI */
	0x00, 0xb0, 0x0d, 0x00, 0x01, 0xc1, 0x00, 0x00,
	/* PAT */
	0x00, 0x01, 0xf0, 0x01,
	/* CRC */
	0x2e, 0x70, 0x19, 0x05,

	/* stuffing 167 bytes */
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,

	/* TS */
	0x47, 0x50, 0x01, 0x10, 0x00,
	/* PSI */
	0x02, 0xb0, 0x17, 0x00, 0x01, 0xc1, 0x00, 0x00,
	/* PMT */
	0xe1, 0x00,
	0xf0, 0x00,
	0x1b, 0xe1, 0x00, 0xf0, 0x00, /* avc epid 256 */
	0x0f, 0xe1, 0x01, 0xf0, 0x00, /* aac epid 257 */
	/* CRC */
	0x2f, 0x44, 0xb9, 0x9b, /* crc for aac */
	/* stuffing 157 bytes */
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
}

// FixedFragmentHeaderHevc 每个TS文件都以固定的PAT，PMT开始
// TODO(chef): [refactor] 考虑去掉这个fixed变量，目前只在测试中使用 [202303]
var FixedFragmentHeaderHevc = []byte{
	/* TS */
	0x47, 0x40, 0x00, 0x10, 0x00,
	/* PSI */
	0x00, 0xb0, 0x0d, 0x00, 0x01, 0xc1, 0x00, 0x00,
	/* PAT */
	0x00, 0x01, 0xf0, 0x01,
	/* CRC */
	0x2e, 0x70, 0x19, 0x05,

	/* stuffing 167 bytes */
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,

	/* TS */
	0x47, 0x50, 0x01, 0x10, 0x00,
	/* PSI */
	0x02, 0xb0, 0x17, 0x00, 0x01, 0xc1, 0x00, 0x00,
	/* PMT */
	0xe1, 0x00,
	0xf0, 0x00,
	//0x1b, 0xe1, 0x00, 0xf0, 0x00, /* avc epid 256 */
	0x24, 0xe1, 0x00, 0xf0, 0x00,
	0x0f, 0xe1, 0x01, 0xf0, 0x00, /* aac epid 257 */
	/* CRC */
	//0x2f, 0x44, 0xb9, 0x9b, /* crc for aac */
	0xc7, 0x72, 0xb7, 0xcb,
	/* stuffing 157 bytes */
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
}

// TS Packet Header
const (
	syncByte uint8 = 0x47

	PidPat   uint16 = 0
	PidPmt   uint16 = 0x1001
	PidVideo uint16 = 0x100
	PidAudio uint16 = 0x101

	// AdaptationFieldControlReserved ------------------------------------------
	// <iso13818-1.pdf> <Table 2-5> <page 38/174>
	// ------------------------------------------
	AdaptationFieldControlReserved uint8 = 0 // Reserved for future use by ISO/IEC
	AdaptationFieldControlNo       uint8 = 1 // No adaptation_field, payload only
	AdaptationFieldControlOnly     uint8 = 2 // Adaptation_field only, no payload
	AdaptationFieldControlFollowed uint8 = 3 // Adaptation_field followed by payload
)

// PMT
const (
	// -----------------------------------------------------------------------------
	// <iso13818-1.pdf> <Table 2-29 Stream type assignments> <page 66/174>
	// 0x0F AAC  (ISO/IEC 13818-7 Audio with ADTS transport syntax)
	// 0x1B AVC  (video stream as defined in ITU-T Rec. H.264 | ISO/IEC 14496-10 Video)
	// 0x24 HEVC (HEVC video stream as defined in Rec. ITU-T H.265 | ISO/IEC 23008-2  MPEG-H Part 2)
	// -----------------------------------------------------------------------------
	StreamTypeUnknown uint8 = 0x00
	StreamTypePrivate uint8 = 0x06
	StreamTypeAac     uint8 = 0x0F
	StreamTypeAvc     uint8 = 0x1B
	StreamTypeHevc    uint8 = 0x24
)

// PES
const (
	// StreamIdAudio -----------------------------------------------------------------
	// <iso13818-1.pdf> <Table 2-18-Stream_id assignments> <page 52/174>
	// -----------------------------------------------------------------
	StreamIdAudio uint8 = 192 // 110x xxxx 0xC0
	StreamIdVideo uint8 = 224 // 1110 xxxx

	// PtsDtsFlags0 ------------------------------
	// <iso13818-1.pdf> <page 53/174>
	// ------------------------------
	PtsDtsFlags0 uint8 = 0 // no PTS no DTS
	PtsDtsFlags1 uint8 = 1 // forbidden
	PtsDtsFlags2 uint8 = 2 // only PTS
	PtsDtsFlags3 uint8 = 3 // both PTS and DTS
)

const (
	delay uint64 = 63000 // 700 ms PCR delay TODO chef: 具体作用？
)
