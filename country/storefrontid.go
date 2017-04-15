package country

import (
	"errors"
	"fmt"
)

var ErrConvertFail = errors.New("convert fail")

// see: https://affiliate.itunes.apple.com/resources/documentation/linking-to-the-itunes-music-store/#appendix
type StorefrontID uint32

const (
	DZ StorefrontID = 143563
	AO StorefrontID = 143564
	AI StorefrontID = 143538
	AG StorefrontID = 143540
	AR StorefrontID = 143505
	AM StorefrontID = 143524
	AU StorefrontID = 143460
	AT StorefrontID = 143445
	AZ StorefrontID = 143568
	BH StorefrontID = 143559
	BD StorefrontID = 143490
	BB StorefrontID = 143541
	BY StorefrontID = 143565
	BE StorefrontID = 143446
	BZ StorefrontID = 143555
	BM StorefrontID = 143542
	BO StorefrontID = 143556
	BW StorefrontID = 143525
	BR StorefrontID = 143503
	VG StorefrontID = 143543
	BN StorefrontID = 143560
	BG StorefrontID = 143526
	CA StorefrontID = 143455
	KY StorefrontID = 143544
	CL StorefrontID = 143483
	CN StorefrontID = 143465
	CO StorefrontID = 143501
	CR StorefrontID = 143495
	CI StorefrontID = 143527
	HR StorefrontID = 143494
	CY StorefrontID = 143557
	CZ StorefrontID = 143489
	DK StorefrontID = 143458
	DM StorefrontID = 143545
	DO StorefrontID = 143508
	EC StorefrontID = 143509
	EG StorefrontID = 143516
	SV StorefrontID = 143506
	EE StorefrontID = 143518
	FI StorefrontID = 143447
	FR StorefrontID = 143442
	DE StorefrontID = 143443
	GH StorefrontID = 143573
	GR StorefrontID = 143448
	GD StorefrontID = 143546
	GT StorefrontID = 143504
	GY StorefrontID = 143553
	HN StorefrontID = 143510
	HK StorefrontID = 143463
	HU StorefrontID = 143482
	IS StorefrontID = 143558
	IN StorefrontID = 143467
	ID StorefrontID = 143476
	IE StorefrontID = 143449
	IL StorefrontID = 143491
	IT StorefrontID = 143450
	JM StorefrontID = 143511
	JP StorefrontID = 143462
	JO StorefrontID = 143528
	KZ StorefrontID = 143517
	KE StorefrontID = 143529
	KR StorefrontID = 143466
	KW StorefrontID = 143493
	LV StorefrontID = 143519
	LB StorefrontID = 143497
	LI StorefrontID = 143522
	LT StorefrontID = 143520
	LU StorefrontID = 143451
	MO StorefrontID = 143515
	MK StorefrontID = 143530
	MG StorefrontID = 143531
	MY StorefrontID = 143473
	MV StorefrontID = 143488
	ML StorefrontID = 143532
	MT StorefrontID = 143521
	MU StorefrontID = 143533
	MX StorefrontID = 143468
	MD StorefrontID = 143523
	MS StorefrontID = 143547
	NP StorefrontID = 143484
	NL StorefrontID = 143452
	NZ StorefrontID = 143461
	NI StorefrontID = 143512
	NE StorefrontID = 143534
	NG StorefrontID = 143561
	NO StorefrontID = 143457
	OM StorefrontID = 143562
	PK StorefrontID = 143477
	PA StorefrontID = 143485
	PY StorefrontID = 143513
	PE StorefrontID = 143507
	PH StorefrontID = 143474
	PL StorefrontID = 143478
	PT StorefrontID = 143453
	QA StorefrontID = 143498
	RO StorefrontID = 143487
	RU StorefrontID = 143469
	SA StorefrontID = 143479
	SN StorefrontID = 143535
	RS StorefrontID = 143500
	SG StorefrontID = 143464
	SK StorefrontID = 143496
	SI StorefrontID = 143499
	ZA StorefrontID = 143472
	ES StorefrontID = 143454
	LK StorefrontID = 143486
	KN StorefrontID = 143548
	LC StorefrontID = 143549
	VC StorefrontID = 143550
	SR StorefrontID = 143554
	SE StorefrontID = 143456
	CH StorefrontID = 143459
	TW StorefrontID = 143470
	TZ StorefrontID = 143572
	TH StorefrontID = 143475
	BS StorefrontID = 143539
	TT StorefrontID = 143551
	TN StorefrontID = 143536
	TR StorefrontID = 143480
	TC StorefrontID = 143552
	UG StorefrontID = 143537
	GB StorefrontID = 143444
	UA StorefrontID = 143492
	AE StorefrontID = 143481
	UY StorefrontID = 143514
	US StorefrontID = 143441
	UZ StorefrontID = 143566
	VE StorefrontID = 143502
	VN StorefrontID = 143471
	YE StorefrontID = 143571
)

func Convert(i uint32) (StorefrontID, error) {
	code := StorefrontID(i).String()
	s := fmt.Sprintf("StorefrontID(%d)", i)
	if code == s {
		return 0, ErrConvertFail
	}
	return StorefrontID(i), nil
}
