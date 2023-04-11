package core

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"
)

type DataSize interface {
	// Convert the DataSize object to the byte representation.
	//
	// Returns: The number of bytes for the data size.
	ToBytes() *float64
	// Convert the DataSize object to its gibibyte representation.
	//
	// If the data size doesn't fit evently into gibibytes it will be rounded
	// up to the closest gibibyte which will be required to hold all the data.
	//
	// Returns: The number of gibibytes for the data size.
	ToGibibytes() *float64
	// Convert the DataSize object to its gigabyte representation.
	//
	// If the data size doesn't fit evently into gigabytes it will be rounded
	// up to the closest gigabyte which will be required to hold all the data.
	//
	// Returns: The number of gigabytes for the data size.
	ToGigabytes() *float64
	// Convert the DataSize object to its kibibyte representation.
	//
	// If the data size doesn't fit evently into kibibytes it will be rounded
	// up to the closest kibibyte which will be required to hold all the data.
	//
	// Returns: The number of kibibytes for the data size.
	ToKibibytes() *float64
	// Convert the DataSize object to its kilobyte representation.
	//
	// If the data size doesn't fit evently into kilobytes it will be rounded
	// up to the closest kilobyte which will be required to hold all the data.
	//
	// Returns: The number of kilobytes for the data size.
	ToKilobytes() *float64
	// Convert the DataSize object to its mebibyte representation.
	//
	// If the data size doesn't fit evently into mebibytes it will be rounded
	// up to the closest mebibyte which will be required to hold all the data.
	//
	// Returns: The number of mebibytes for the data size.
	ToMebibytes() *float64
	// Convert the DataSize object to its megabyte representation.
	//
	// If the data size doesn't fit evently into megabytes it will be rounded
	// up to the closest megabyte which will be required to hold all the data.
	//
	// Returns: The number of megabytes for the data size.
	ToMegabytes() *float64
	// Convert the DataSize object to its pebibyte representation.
	//
	// If the data size doesn't fit evently into pebibytes it will be rounded
	// up to the closest pebibyte which will be required to hold all the data.
	//
	// Returns: The number of pebibytes for the data size.
	ToPebibytes() *float64
	// Convert the DataSize object to its petabyte representation.
	//
	// If the data size doesn't fit evently into petabytes it will be rounded
	// up to the closest petabyte which will be required to hold all the data.
	//
	// Returns: The number of petabytes for the data size.
	ToPetabytes() *float64
	// Convert the DataSize object to its tebibyte representation.
	//
	// If the data size doesn't fit evently into tebibytes it will be rounded
	// up to the closest tebibyte which will be required to hold all the data.
	//
	// Returns: The number of tebibytes for the data size.
	ToTebibytes() *float64
	// Convert the DataSize object to its terabyte representation.
	//
	// If the data size doesn't fit evently into terabytes it will be rounded
	// up to the closest terabyte which will be required to hold all the data.
	//
	// Returns: The number of terabytes for the data size.
	ToTerabytes() *float64
}

// The jsii proxy struct for DataSize
type jsiiProxy_DataSize struct {
	_ byte // padding
}

// Create a `DataSize` representing an amount of bytes.
//
// Returns: A `DataSize` representing the specified number of bytes.
func DataSize_Bytes(bytes *float64) DataSize {
	_init_.Initialize()

	if err := validateDataSize_BytesParameters(bytes); err != nil {
		panic(err)
	}
	var returns DataSize

	_jsii_.StaticInvoke(
		"cdk-extensions.core.DataSize",
		"bytes",
		[]interface{}{bytes},
		&returns,
	)

	return returns
}

// Create a `DataSize` representing an amount of gibibytes.
//
// Returns: A `DataSize` representing the specified number of gibibytes.
func DataSize_Gibibytes(gibibytes *float64) DataSize {
	_init_.Initialize()

	if err := validateDataSize_GibibytesParameters(gibibytes); err != nil {
		panic(err)
	}
	var returns DataSize

	_jsii_.StaticInvoke(
		"cdk-extensions.core.DataSize",
		"gibibytes",
		[]interface{}{gibibytes},
		&returns,
	)

	return returns
}

// Create a `DataSize` representing an amount of gigabytes.
//
// Returns: A `DataSize` representing the specified number of gigabytes.
func DataSize_Gigabytes(gigabytes *float64) DataSize {
	_init_.Initialize()

	if err := validateDataSize_GigabytesParameters(gigabytes); err != nil {
		panic(err)
	}
	var returns DataSize

	_jsii_.StaticInvoke(
		"cdk-extensions.core.DataSize",
		"gigabytes",
		[]interface{}{gigabytes},
		&returns,
	)

	return returns
}

// Create a `DataSize` representing an amount of kibibytes.
//
// Returns: A `DataSize` representing the specified number of kibibytes.
func DataSize_Kibibytes(kibibytes *float64) DataSize {
	_init_.Initialize()

	if err := validateDataSize_KibibytesParameters(kibibytes); err != nil {
		panic(err)
	}
	var returns DataSize

	_jsii_.StaticInvoke(
		"cdk-extensions.core.DataSize",
		"kibibytes",
		[]interface{}{kibibytes},
		&returns,
	)

	return returns
}

// Create a `DataSize` representing an amount of kilobytes.
//
// Returns: A `DataSize` representing the specified number of kilobytes.
func DataSize_Kilobytes(kilobytes *float64) DataSize {
	_init_.Initialize()

	if err := validateDataSize_KilobytesParameters(kilobytes); err != nil {
		panic(err)
	}
	var returns DataSize

	_jsii_.StaticInvoke(
		"cdk-extensions.core.DataSize",
		"kilobytes",
		[]interface{}{kilobytes},
		&returns,
	)

	return returns
}

// Create a `DataSize` representing an amount of mebibytes.
//
// Returns: A `DataSize` representing the specified number of mebibytes.
func DataSize_Mebibytes(mebibytes *float64) DataSize {
	_init_.Initialize()

	if err := validateDataSize_MebibytesParameters(mebibytes); err != nil {
		panic(err)
	}
	var returns DataSize

	_jsii_.StaticInvoke(
		"cdk-extensions.core.DataSize",
		"mebibytes",
		[]interface{}{mebibytes},
		&returns,
	)

	return returns
}

// Create a `DataSize` representing an amount of megabytes.
//
// Returns: A `DataSize` representing the specified number of megabytes.
func DataSize_Megabytes(megabytes *float64) DataSize {
	_init_.Initialize()

	if err := validateDataSize_MegabytesParameters(megabytes); err != nil {
		panic(err)
	}
	var returns DataSize

	_jsii_.StaticInvoke(
		"cdk-extensions.core.DataSize",
		"megabytes",
		[]interface{}{megabytes},
		&returns,
	)

	return returns
}

// Create a `DataSize` representing an amount of pebibytes.
//
// Returns: A `DataSize` representing the specified number of pebibytes.
func DataSize_Pebibytes(pebibytes *float64) DataSize {
	_init_.Initialize()

	if err := validateDataSize_PebibytesParameters(pebibytes); err != nil {
		panic(err)
	}
	var returns DataSize

	_jsii_.StaticInvoke(
		"cdk-extensions.core.DataSize",
		"pebibytes",
		[]interface{}{pebibytes},
		&returns,
	)

	return returns
}

// Create a `DataSize` representing an amount of petabytes.
//
// Returns: A `DataSize` representing the specified number of petabytes.
func DataSize_Petabytes(petabytes *float64) DataSize {
	_init_.Initialize()

	if err := validateDataSize_PetabytesParameters(petabytes); err != nil {
		panic(err)
	}
	var returns DataSize

	_jsii_.StaticInvoke(
		"cdk-extensions.core.DataSize",
		"petabytes",
		[]interface{}{petabytes},
		&returns,
	)

	return returns
}

// Create a `DataSize` representing an amount of tebibytes.
//
// Returns: A `DataSize` representing the specified number of tebibytes.
func DataSize_Tebibytes(tebibytes *float64) DataSize {
	_init_.Initialize()

	if err := validateDataSize_TebibytesParameters(tebibytes); err != nil {
		panic(err)
	}
	var returns DataSize

	_jsii_.StaticInvoke(
		"cdk-extensions.core.DataSize",
		"tebibytes",
		[]interface{}{tebibytes},
		&returns,
	)

	return returns
}

// Create a `DataSize` representing an amount of terabytes.
//
// Returns: A `DataSize` representing the specified number of terabytes.
func DataSize_Terabytes(terabytes *float64) DataSize {
	_init_.Initialize()

	if err := validateDataSize_TerabytesParameters(terabytes); err != nil {
		panic(err)
	}
	var returns DataSize

	_jsii_.StaticInvoke(
		"cdk-extensions.core.DataSize",
		"terabytes",
		[]interface{}{terabytes},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSize) ToBytes() *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"toBytes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSize) ToGibibytes() *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"toGibibytes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSize) ToGigabytes() *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"toGigabytes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSize) ToKibibytes() *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"toKibibytes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSize) ToKilobytes() *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"toKilobytes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSize) ToMebibytes() *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"toMebibytes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSize) ToMegabytes() *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"toMegabytes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSize) ToPebibytes() *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"toPebibytes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSize) ToPetabytes() *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"toPetabytes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSize) ToTebibytes() *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"toTebibytes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSize) ToTerabytes() *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"toTerabytes",
		nil, // no parameters
		&returns,
	)

	return returns
}

