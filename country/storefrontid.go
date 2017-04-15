package country

import (
	"errors"
	"fmt"
)

var ErrConvertFail = errors.New("convert fail")

// see: https://affiliate.itunes.apple.com/resources/documentation/linking-to-the-itunes-music-store/#appendix
type StorefrontCode uint32

const (
	DZ StorefrontCode = 143563
	AO StorefrontCode = 143564
	AI StorefrontCode = 143538
	AG StorefrontCode = 143540
	AR StorefrontCode = 143505
	AM StorefrontCode = 143524
	AU StorefrontCode = 143460
	AT StorefrontCode = 143445
	AZ StorefrontCode = 143568
	BH StorefrontCode = 143559
	BD StorefrontCode = 143490
	BB StorefrontCode = 143541
	BY StorefrontCode = 143565
	BE StorefrontCode = 143446
	BZ StorefrontCode = 143555
	BM StorefrontCode = 143542
	BO StorefrontCode = 143556
	BW StorefrontCode = 143525
	BR StorefrontCode = 143503
	VG StorefrontCode = 143543
	BN StorefrontCode = 143560
	BG StorefrontCode = 143526
	CA StorefrontCode = 143455
	KY StorefrontCode = 143544
	CL StorefrontCode = 143483
	CN StorefrontCode = 143465
	CO StorefrontCode = 143501
	CR StorefrontCode = 143495
	CI StorefrontCode = 143527
	HR StorefrontCode = 143494
	CY StorefrontCode = 143557
	CZ StorefrontCode = 143489
	DK StorefrontCode = 143458
	DM StorefrontCode = 143545
	DO StorefrontCode = 143508
	EC StorefrontCode = 143509
	EG StorefrontCode = 143516
	SV StorefrontCode = 143506
	EE StorefrontCode = 143518
	FI StorefrontCode = 143447
	FR StorefrontCode = 143442
	DE StorefrontCode = 143443
	GH StorefrontCode = 143573
	GR StorefrontCode = 143448
	GD StorefrontCode = 143546
	GT StorefrontCode = 143504
	GY StorefrontCode = 143553
	HN StorefrontCode = 143510
	HK StorefrontCode = 143463
	HU StorefrontCode = 143482
	IS StorefrontCode = 143558
	IN StorefrontCode = 143467
	ID StorefrontCode = 143476
	IE StorefrontCode = 143449
	IL StorefrontCode = 143491
	IT StorefrontCode = 143450
	JM StorefrontCode = 143511
	JP StorefrontCode = 143462
	JO StorefrontCode = 143528
	KZ StorefrontCode = 143517
	KE StorefrontCode = 143529
	KR StorefrontCode = 143466
	KW StorefrontCode = 143493
	LV StorefrontCode = 143519
	LB StorefrontCode = 143497
	LI StorefrontCode = 143522
	LT StorefrontCode = 143520
	LU StorefrontCode = 143451
	MO StorefrontCode = 143515
	MK StorefrontCode = 143530
	MG StorefrontCode = 143531
	MY StorefrontCode = 143473
	MV StorefrontCode = 143488
	ML StorefrontCode = 143532
	MT StorefrontCode = 143521
	MU StorefrontCode = 143533
	MX StorefrontCode = 143468
	MD StorefrontCode = 143523
	MS StorefrontCode = 143547
	NP StorefrontCode = 143484
	NL StorefrontCode = 143452
	NZ StorefrontCode = 143461
	NI StorefrontCode = 143512
	NE StorefrontCode = 143534
	NG StorefrontCode = 143561
	NO StorefrontCode = 143457
	OM StorefrontCode = 143562
	PK StorefrontCode = 143477
	PA StorefrontCode = 143485
	PY StorefrontCode = 143513
	PE StorefrontCode = 143507
	PH StorefrontCode = 143474
	PL StorefrontCode = 143478
	PT StorefrontCode = 143453
	QA StorefrontCode = 143498
	RO StorefrontCode = 143487
	RU StorefrontCode = 143469
	SA StorefrontCode = 143479
	SN StorefrontCode = 143535
	RS StorefrontCode = 143500
	SG StorefrontCode = 143464
	SK StorefrontCode = 143496
	SI StorefrontCode = 143499
	ZA StorefrontCode = 143472
	ES StorefrontCode = 143454
	LK StorefrontCode = 143486
	KN StorefrontCode = 143548
	LC StorefrontCode = 143549
	VC StorefrontCode = 143550
	SR StorefrontCode = 143554
	SE StorefrontCode = 143456
	CH StorefrontCode = 143459
	TW StorefrontCode = 143470
	TZ StorefrontCode = 143572
	TH StorefrontCode = 143475
	BS StorefrontCode = 143539
	TT StorefrontCode = 143551
	TN StorefrontCode = 143536
	TR StorefrontCode = 143480
	TC StorefrontCode = 143552
	UG StorefrontCode = 143537
	GB StorefrontCode = 143444
	UA StorefrontCode = 143492
	AE StorefrontCode = 143481
	UY StorefrontCode = 143514
	US StorefrontCode = 143441
	UZ StorefrontCode = 143566
	VE StorefrontCode = 143502
	VN StorefrontCode = 143471
	YE StorefrontCode = 143571
)

func Convert(i uint32) (StorefrontCode, error) {
	code := StorefrontCode(i).String()
	s := fmt.Sprintf("StorefrontCode(%d)", i)
	if code == s {
		return 0, ErrConvertFail
	}
	return StorefrontCode(i), nil
}
