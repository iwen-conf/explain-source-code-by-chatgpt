// Copyright 2016 - 2023 The excelize Authors. All rights reserved. Use of
// this source code is governed by a BSD-style license that can be found in
// the LICENSE file.
//
// Package excelize providing a set of functions that allow you to write to and
// read from XLAM / XLSM / XLSX / XLTM / XLTX files. Supports reading and
// writing spreadsheet documents generated by Microsoft Excel™ 2007 and later.
// Supports complex components by high compatibility, and provided streaming
// API for generating or reading data from a worksheet with huge amounts of
// data. This library needs Go version 1.16 or later.
//
// This file contains default templates for XML files we don't yet populated
// based on content.

package excelize

import "encoding/xml"

// Source relationship and namespace list, associated prefixes and schema in which it was
// introduced.
var (
	NameSpaceDocumentPropertiesVariantTypes = xml.Attr{Name: xml.Name{Local: "vt", Space: "xmlns"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes"}
	NameSpaceDrawing2016SVG                 = xml.Attr{Name: xml.Name{Local: "asvg", Space: "xmlns"}, Value: "http://schemas.microsoft.com/office/drawing/2016/SVG/main"}
	NameSpaceDrawingML                      = xml.Attr{Name: xml.Name{Local: "a", Space: "xmlns"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"}
	NameSpaceDrawingMLA14                   = xml.Attr{Name: xml.Name{Local: "a14", Space: "xmlns"}, Value: "http://schemas.microsoft.com/office/drawing/2010/main"}
	NameSpaceDrawingMLChart                 = xml.Attr{Name: xml.Name{Local: "c", Space: "xmlns"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/chart"}
	NameSpaceDrawingMLSlicer                = xml.Attr{Name: xml.Name{Local: "sle", Space: "xmlns"}, Value: "http://schemas.microsoft.com/office/drawing/2010/slicer"}
	NameSpaceDrawingMLSlicerX15             = xml.Attr{Name: xml.Name{Local: "sle15", Space: "xmlns"}, Value: "http://schemas.microsoft.com/office/drawing/2012/slicer"}
	NameSpaceDrawingMLSpreadSheet           = xml.Attr{Name: xml.Name{Local: "xdr", Space: "xmlns"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"}
	NameSpaceMacExcel2008Main               = xml.Attr{Name: xml.Name{Local: "mx", Space: "xmlns"}, Value: "http://schemas.microsoft.com/office/mac/excel/2008/main"}
	NameSpaceSpreadSheet                    = xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"}
	NameSpaceSpreadSheetExcel2006Main       = xml.Attr{Name: xml.Name{Local: "xne", Space: "xmlns"}, Value: "http://schemas.microsoft.com/office/excel/2006/main"}
	NameSpaceSpreadSheetX14                 = xml.Attr{Name: xml.Name{Local: "x14", Space: "xmlns"}, Value: "http://schemas.microsoft.com/office/spreadsheetml/2009/9/main"}
	NameSpaceSpreadSheetX15                 = xml.Attr{Name: xml.Name{Local: "x15", Space: "xmlns"}, Value: "http://schemas.microsoft.com/office/spreadsheetml/2010/11/main"}
	NameSpaceSpreadSheetXR10                = xml.Attr{Name: xml.Name{Local: "xr10", Space: "xmlns"}, Value: "http://schemas.microsoft.com/office/spreadsheetml/2016/revision10"}
	SourceRelationship                      = xml.Attr{Name: xml.Name{Local: "r", Space: "xmlns"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"}
	SourceRelationshipChart20070802         = xml.Attr{Name: xml.Name{Local: "c14", Space: "xmlns"}, Value: "http://schemas.microsoft.com/office/drawing/2007/8/2/chart"}
	SourceRelationshipChart2014             = xml.Attr{Name: xml.Name{Local: "c16", Space: "xmlns"}, Value: "http://schemas.microsoft.com/office/drawing/2014/chart"}
	SourceRelationshipChart201506           = xml.Attr{Name: xml.Name{Local: "c16r2", Space: "xmlns"}, Value: "http://schemas.microsoft.com/office/drawing/2015/06/chart"}
	SourceRelationshipCompatibility         = xml.Attr{Name: xml.Name{Local: "mc", Space: "xmlns"}, Value: "http://schemas.openxmlformats.org/markup-compatibility/2006"}
)

// Source relationship and namespace.
const (
	ContentTypeAddinMacro                         = "application/vnd.ms-excel.addin.macroEnabled.main+xml"
	ContentTypeDrawing                            = "application/vnd.openxmlformats-officedocument.drawing+xml"
	ContentTypeDrawingML                          = "application/vnd.openxmlformats-officedocument.drawingml.chart+xml"
	ContentTypeMacro                              = "application/vnd.ms-excel.sheet.macroEnabled.main+xml"
	ContentTypeRelationships                      = "application/vnd.openxmlformats-package.relationships+xml"
	ContentTypeSheetML                            = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet.main+xml"
	ContentTypeSlicer                             = "application/vnd.ms-excel.slicer+xml"
	ContentTypeSlicerCache                        = "application/vnd.ms-excel.slicerCache+xml"
	ContentTypeSpreadSheetMLChartsheet            = "application/vnd.openxmlformats-officedocument.spreadsheetml.chartsheet+xml"
	ContentTypeSpreadSheetMLComments              = "application/vnd.openxmlformats-officedocument.spreadsheetml.comments+xml"
	ContentTypeSpreadSheetMLPivotCacheDefinition  = "application/vnd.openxmlformats-officedocument.spreadsheetml.pivotCacheDefinition+xml"
	ContentTypeSpreadSheetMLPivotTable            = "application/vnd.openxmlformats-officedocument.spreadsheetml.pivotTable+xml"
	ContentTypeSpreadSheetMLSharedStrings         = "application/vnd.openxmlformats-officedocument.spreadsheetml.sharedStrings+xml"
	ContentTypeSpreadSheetMLTable                 = "application/vnd.openxmlformats-officedocument.spreadsheetml.table+xml"
	ContentTypeSpreadSheetMLWorksheet             = "application/vnd.openxmlformats-officedocument.spreadsheetml.worksheet+xml"
	ContentTypeTemplate                           = "application/vnd.openxmlformats-officedocument.spreadsheetml.template.main+xml"
	ContentTypeTemplateMacro                      = "application/vnd.ms-excel.template.macroEnabled.main+xml"
	ContentTypeVBA                                = "application/vnd.ms-office.vbaProject"
	ContentTypeVML                                = "application/vnd.openxmlformats-officedocument.vmlDrawing"
	NameSpaceDrawingMLMain                        = "http://schemas.openxmlformats.org/drawingml/2006/main"
	NameSpaceDublinCore                           = "http://purl.org/dc/elements/1.1/"
	NameSpaceDublinCoreMetadataInitiative         = "http://purl.org/dc/dcmitype/"
	NameSpaceDublinCoreTerms                      = "http://purl.org/dc/terms/"
	NameSpaceExtendedProperties                   = "http://schemas.openxmlformats.org/officeDocument/2006/extended-properties"
	NameSpaceXML                                  = "http://www.w3.org/XML/1998/namespace"
	NameSpaceXMLSchemaInstance                    = "http://www.w3.org/2001/XMLSchema-instance"
	SourceRelationshipChart                       = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/chart"
	SourceRelationshipChartsheet                  = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/chartsheet"
	SourceRelationshipComments                    = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/comments"
	SourceRelationshipDialogsheet                 = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/dialogsheet"
	SourceRelationshipDrawingML                   = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/drawing"
	SourceRelationshipDrawingVML                  = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/vmlDrawing"
	SourceRelationshipExtendProperties            = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/extended-properties"
	SourceRelationshipHyperLink                   = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/hyperlink"
	SourceRelationshipImage                       = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/image"
	SourceRelationshipOfficeDocument              = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument"
	SourceRelationshipPivotCache                  = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/pivotCacheDefinition"
	SourceRelationshipPivotTable                  = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/pivotTable"
	SourceRelationshipSharedStrings               = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/sharedStrings"
	SourceRelationshipSlicer                      = "http://schemas.microsoft.com/office/2007/relationships/slicer"
	SourceRelationshipSlicerCache                 = "http://schemas.microsoft.com/office/2007/relationships/slicerCache"
	SourceRelationshipTable                       = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/table"
	SourceRelationshipVBAProject                  = "http://schemas.microsoft.com/office/2006/relationships/vbaProject"
	SourceRelationshipWorkSheet                   = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/worksheet"
	StrictNameSpaceDocumentPropertiesVariantTypes = "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes"
	StrictNameSpaceDrawingMLMain                  = "http://purl.oclc.org/ooxml/drawingml/main"
	StrictNameSpaceExtendedProperties             = "http://purl.oclc.org/ooxml/officeDocument/extendedProperties"
	StrictNameSpaceSpreadSheet                    = "http://purl.oclc.org/ooxml/spreadsheetml/main"
	StrictSourceRelationship                      = "http://purl.oclc.org/ooxml/officeDocument/relationships"
	StrictSourceRelationshipChart                 = "http://purl.oclc.org/ooxml/officeDocument/relationships/chart"
	StrictSourceRelationshipComments              = "http://purl.oclc.org/ooxml/officeDocument/relationships/comments"
	StrictSourceRelationshipExtendProperties      = "http://purl.oclc.org/ooxml/officeDocument/relationships/extendedProperties"
	StrictSourceRelationshipImage                 = "http://purl.oclc.org/ooxml/officeDocument/relationships/image"
	StrictSourceRelationshipOfficeDocument        = "http://purl.oclc.org/ooxml/officeDocument/relationships/officeDocument"
	// The following constants defined the extLst child element
	// ([ISO/IEC29500-1:2016] section 18.2.10) of the workbook and worksheet
	// elements extended by the addition of new child ext elements.
	ExtURICalcFeatures                   = "{B58B0392-4F1F-4190-BB64-5DF3571DCE5F}"
	ExtURIConditionalFormattingRuleID    = "{B025F937-C7B1-47D3-B67F-A62EFF666E3E}"
	ExtURIConditionalFormattings         = "{78C0D931-6437-407d-A8EE-F0AAD7539E65}"
	ExtURIDataModel                      = "{FCE2AD5D-F65C-4FA6-A056-5C36A1767C68}"
	ExtURIDataValidations                = "{CCE6A557-97BC-4B89-ADB6-D9C93CAAB3DF}"
	ExtURIDrawingBlip                    = "{28A0092B-C50C-407E-A947-70E740481C1C}"
	ExtURIExternalLinkPr                 = "{FCE6A71B-6B00-49CD-AB44-F6B1AE7CDE65}"
	ExtURIIgnoredErrors                  = "{01252117-D84E-4E92-8308-4BE1C098FCBB}"
	ExtURIMacExcelMX                     = "{64002731-A6B0-56B0-2670-7721B7C09600}"
	ExtURIModelTimeGroupings             = "{9835A34E-60A6-4A7C-AAB8-D5F71C897F49}"
	ExtURIPivotCacheDefinition           = "{725AE2AE-9491-48be-B2B4-4EB974FC3084}"
	ExtURIPivotCachesX14                 = "{876F7934-8845-4945-9796-88D515C7AA90}"
	ExtURIPivotCachesX15                 = "{841E416B-1EF1-43b6-AB56-02D37102CBD5}"
	ExtURIPivotTableReferences           = "{983426D0-5260-488c-9760-48F4B6AC55F4}"
	ExtURIProtectedRanges                = "{FC87AEE6-9EDD-4A0A-B7FB-166176984837}"
	ExtURISlicerCacheDefinition          = "{2F2917AC-EB37-4324-AD4E-5DD8C200BD13}"
	ExtURISlicerCacheHideItemsWithNoData = "{470722E0-AACD-4C17-9CDC-17EF765DBC7E}"
	ExtURISlicerCachesX14                = "{BBE1A952-AA13-448e-AADC-164F8A28A991}"
	ExtURISlicerCachesX15                = "{46BE6895-7355-4a93-B00E-2C351335B9C9}"
	ExtURISlicerListX14                  = "{A8765BA9-456A-4dab-B4F3-ACF838C121DE}"
	ExtURISlicerListX15                  = "{3A4CF648-6AED-40f4-86FF-DC5316D8AED3}"
	ExtURISparklineGroups                = "{05C60535-1F16-4fd2-B633-F4F36F0B64E0}"
	ExtURISVG                            = "{96DAC541-7B7A-43D3-8B79-37D633B846F1}"
	ExtURITimelineCachePivotCaches       = "{A2CB5862-8E78-49c6-8D9D-AF26E26ADB89}"
	ExtURITimelineCacheRefs              = "{D0CA8CA8-9F24-4464-BF8E-62219DCF47F9}"
	ExtURITimelineRefs                   = "{7E03D99C-DC04-49d9-9315-930204A7B6E9}"
	ExtURIWebExtensions                  = "{F7C9EE02-42E1-4005-9D12-6889AFFD525C}"
	ExtURIWorkbookPrX14                  = "{79F54976-1DA5-4618-B147-ACDE4B953A38}"
	ExtURIWorkbookPrX15                  = "{140A7094-0E35-4892-8432-C4D2E57EDEB5}"
)

// workbookExtURIPriority is the priority of URI in the workbook extension lists.
var workbookExtURIPriority = []string{
	ExtURIPivotCachesX14,
	ExtURISlicerCachesX14,
	ExtURISlicerCachesX15,
	ExtURIWorkbookPrX14,
	ExtURIPivotCachesX15,
	ExtURIPivotTableReferences,
	ExtURITimelineCachePivotCaches,
	ExtURITimelineCacheRefs,
	ExtURIWorkbookPrX15,
	ExtURIDataModel,
	ExtURICalcFeatures,
	ExtURIExternalLinkPr,
	ExtURIModelTimeGroupings,
}

// worksheetExtURIPriority is the priority of URI in the worksheet extension lists.
var worksheetExtURIPriority = []string{
	ExtURIConditionalFormattings,
	ExtURIDataValidations,
	ExtURISparklineGroups,
	ExtURISlicerListX14,
	ExtURIProtectedRanges,
	ExtURIIgnoredErrors,
	ExtURIWebExtensions,
	ExtURISlicerListX15,
	ExtURITimelineRefs,
	ExtURIExternalLinkPr,
}

// Excel specifications and limits
const (
	MaxCellStyles        = 65430
	MaxColumns           = 16384
	MaxColumnWidth       = 255
	MaxFieldLength       = 255
	MaxFilePathLength    = 207
	MaxFormControlValue  = 30000
	MaxFontFamilyLength  = 31
	MaxFontSize          = 409
	MaxRowHeight         = 409
	MaxSheetNameLength   = 31
	MinColumns           = 1
	MinFontSize          = 1
	StreamChunkSize      = 1 << 24
	TotalCellChars       = 32767
	TotalRows            = 1048576
	TotalSheetHyperlinks = 65529
	UnzipSizeLimit       = 1000 << 24
	// pivotTableVersion should be greater than 3. One or more of the
	// PivotTables chosen are created in a version of Excel earlier than
	// Excel 2007 or in compatibility mode. Slicer can only be used with
	// PivotTables created in Excel 2007 or a newer version of Excel.
	pivotTableVersion           = 3
	pivotTableRefreshedVersion  = 8
	defaultDrawingScale         = 1.0
	defaultChartDimensionWidth  = 480
	defaultChartDimensionHeight = 260
	defaultSlicerWidth          = 200
	defaultSlicerHeight         = 200
	defaultChartLegendPosition  = "bottom"
	defaultChartShowBlanksAs    = "gap"
	defaultShapeSize            = 160
	defaultShapeLineWidth       = 1
)

// ColorMappingType is the type of color transformation.
type ColorMappingType byte

// Color transformation types enumeration.
const (
	ColorMappingTypeLight1 ColorMappingType = iota
	ColorMappingTypeDark1
	ColorMappingTypeLight2
	ColorMappingTypeDark2
	ColorMappingTypeAccent1
	ColorMappingTypeAccent2
	ColorMappingTypeAccent3
	ColorMappingTypeAccent4
	ColorMappingTypeAccent5
	ColorMappingTypeAccent6
	ColorMappingTypeHyperlink
	ColorMappingTypeFollowedHyperlink
	ColorMappingTypeUnset int = -1
)

const (
	defaultTempFileSST          = "sharedStrings"
	defaultXMLPathCalcChain     = "xl/calcChain.xml"
	defaultXMLPathContentTypes  = "[Content_Types].xml"
	defaultXMLPathDocPropsApp   = "docProps/app.xml"
	defaultXMLPathDocPropsCore  = "docProps/core.xml"
	defaultXMLPathSharedStrings = "xl/sharedStrings.xml"
	defaultXMLPathStyles        = "xl/styles.xml"
	defaultXMLPathTheme         = "xl/theme/theme1.xml"
	defaultXMLPathVolatileDeps  = "xl/volatileDependencies.xml"
	defaultXMLPathWorkbook      = "xl/workbook.xml"
	defaultXMLPathWorkbookRels  = "xl/_rels/workbook.xml.rels"
)

// IndexedColorMapping is the table of default mappings from indexed color value
// to RGB value. Note that 0-7 are redundant of 8-15 to preserve backwards
// compatibility. A legacy indexing scheme for colors that is still required
// for some records, and for backwards compatibility with legacy formats. This
// element contains a sequence of RGB color values that correspond to color
// indexes (zero-based). When using the default indexed color palette, the
// values are not written out, but instead are implied. When the color palette
// has been modified from default, then the entire color palette is written
// out.
var IndexedColorMapping = []string{
	"000000", "FFFFFF", "FF0000", "00FF00", "0000FF", "FFFF00", "FF00FF", "00FFFF",
	"000000", "FFFFFF", "FF0000", "00FF00", "0000FF", "FFFF00", "FF00FF", "00FFFF",
	"800000", "008000", "000080", "808000", "800080", "008080", "C0C0C0", "808080",
	"9999FF", "993366", "FFFFCC", "CCFFFF", "660066", "FF8080", "0066CC", "CCCCFF",
	"000080", "FF00FF", "FFFF00", "00FFFF", "800080", "800000", "008080", "0000FF",
	"00CCFF", "CCFFFF", "CCFFCC", "FFFF99", "99CCFF", "FF99CC", "CC99FF", "FFCC99",
	"3366FF", "33CCCC", "99CC00", "FFCC00", "FF9900", "FF6600", "666699", "969696",
	"003366", "339966", "003300", "333300", "993300", "993366", "333399", "333333",
	"000000", "FFFFFF",
}

// supportedImageTypes defined supported image types.
var supportedImageTypes = map[string]string{
	".bmp": ".bmp", ".emf": ".emf", ".emz": ".emz", ".gif": ".gif",
	".jpeg": ".jpeg", ".jpg": ".jpeg", ".png": ".png", ".svg": ".svg",
	".tif": ".tiff", ".tiff": ".tiff", ".wmf": ".wmf", ".wmz": ".wmz",
}

// supportedContentTypes defined supported file format types.
var supportedContentTypes = map[string]string{
	".xlam": ContentTypeAddinMacro,
	".xlsm": ContentTypeMacro,
	".xlsx": ContentTypeSheetML,
	".xltm": ContentTypeTemplateMacro,
	".xltx": ContentTypeTemplate,
}

// supportedUnderlineTypes defined supported underline types.
var supportedUnderlineTypes = []string{"none", "single", "double"}

// supportedDrawingUnderlineTypes defined supported underline types in drawing
// markup language.
var supportedDrawingUnderlineTypes = []string{
	"none", "words", "sng", "dbl", "heavy", "dotted", "dottedHeavy", "dash", "dashHeavy", "dashLong", "dashLongHeavy", "dotDash", "dotDashHeavy", "dotDotDash", "dotDotDashHeavy", "wavy", "wavyHeavy",
	"wavyDbl",
}

// supportedPositioning defined supported positioning types.
var supportedPositioning = []string{"absolute", "oneCell", "twoCell"}

// builtInDefinedNames defined built-in defined names are built with a _xlnm prefix.
var builtInDefinedNames = []string{"_xlnm.Print_Area", "_xlnm.Print_Titles", "_xlnm._FilterDatabase"}

const templateDocpropsApp = `<Properties xmlns="http://schemas.openxmlformats.org/officeDocument/2006/extended-properties" xmlns:vt="http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes"><TotalTime>0</TotalTime><Application>Go Excelize</Application></Properties>`

const templateContentTypes = `<Types xmlns="http://schemas.openxmlformats.org/package/2006/content-types"><Override PartName="/xl/theme/theme1.xml" ContentType="application/vnd.openxmlformats-officedocument.theme+xml"/><Override PartName="/xl/styles.xml" ContentType="application/vnd.openxmlformats-officedocument.spreadsheetml.styles+xml"/><Default Extension="rels" ContentType="application/vnd.openxmlformats-package.relationships+xml"/><Default Extension="xml" ContentType="application/xml"/><Override PartName="/xl/workbook.xml" ContentType="application/vnd.openxmlformats-officedocument.spreadsheetml.sheet.main+xml"/><Override PartName="/docProps/app.xml" ContentType="application/vnd.openxmlformats-officedocument.extended-properties+xml"/><Override PartName="/xl/worksheets/sheet1.xml" ContentType="application/vnd.openxmlformats-officedocument.spreadsheetml.worksheet+xml"/><Override PartName="/docProps/core.xml" ContentType="application/vnd.openxmlformats-package.core-properties+xml"/></Types>`

const templateWorkbook = `<workbook xmlns="http://schemas.openxmlformats.org/spreadsheetml/2006/main" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006" mc:Ignorable="x15" xmlns:x15="http://schemas.microsoft.com/office/spreadsheetml/2010/11/main"><fileVersion appName="xl" lastEdited="6" lowestEdited="6" rupBuild="14420" /><workbookPr filterPrivacy="1" defaultThemeVersion="164011" /><bookViews><workbookView xWindow="0" yWindow="0" windowWidth="14805" windowHeight="8010" /></bookViews><sheets><sheet name="Sheet1" sheetId="1" r:id="rId1" /></sheets><calcPr calcId="122211" /></workbook>`

const templateStyles = `<styleSheet xmlns="http://schemas.openxmlformats.org/spreadsheetml/2006/main" xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006" mc:Ignorable="x14ac x16r2" xmlns:x14ac="http://schemas.microsoft.com/office/spreadsheetml/2009/9/ac" xmlns:x16r2="http://schemas.microsoft.com/office/spreadsheetml/2015/02/main"><fonts count="1" x14ac:knownFonts="1"><font><sz val="11"/><color theme="1"/><name val="Calibri"/><family val="2"/></font></fonts><fills count="2"><fill><patternFill patternType="none"/></fill><fill><patternFill patternType="gray125"/></fill></fills><borders count="1"><border><left/><right/><top/><bottom/><diagonal/></border></borders><cellStyleXfs count="1"><xf numFmtId="0" fontId="0" fillId="0" borderId="0"/></cellStyleXfs><cellXfs count="1"><xf numFmtId="0" fontId="0" fillId="0" borderId="0" xfId="0"/></cellXfs><cellStyles count="1"><cellStyle name="Normal" xfId="0" builtinId="0"/></cellStyles><dxfs count="0"/><tableStyles count="0" defaultTableStyle="TableStyleMedium2" defaultPivotStyle="PivotStyleLight16"/></styleSheet>`

const templateSheet = `<worksheet xmlns="http://schemas.openxmlformats.org/spreadsheetml/2006/main" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships"><dimension ref="A1"/><sheetViews><sheetView tabSelected="1" workbookViewId="0"/></sheetViews><sheetFormatPr defaultRowHeight="15"/><sheetData/></worksheet>`

const templateWorkbookRels = `<Relationships xmlns="http://schemas.openxmlformats.org/package/2006/relationships"><Relationship Id="rId1" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/worksheet" Target="worksheets/sheet1.xml"/><Relationship Id="rId2" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/styles" Target="styles.xml"/><Relationship Id="rId3" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/theme" Target="theme/theme1.xml"/></Relationships>`

const templateDocpropsCore = `<cp:coreProperties xmlns:cp="http://schemas.openxmlformats.org/package/2006/metadata/core-properties" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:dcterms="http://purl.org/dc/terms/" xmlns:dcmitype="http://purl.org/dc/dcmitype/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><dc:creator>xuri</dc:creator><dcterms:created xsi:type="dcterms:W3CDTF">2006-09-16T00:00:00Z</dcterms:created><dcterms:modified xsi:type="dcterms:W3CDTF">2006-09-16T00:00:00Z</dcterms:modified></cp:coreProperties>`

const templateRels = `<Relationships xmlns="http://schemas.openxmlformats.org/package/2006/relationships"><Relationship Id="rId3" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/extended-properties" Target="docProps/app.xml"/><Relationship Id="rId2" Type="http://schemas.openxmlformats.org/package/2006/relationships/metadata/core-properties" Target="docProps/core.xml"/><Relationship Id="rId1" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument" Target="xl/workbook.xml"/></Relationships>`

const templateTheme = `<a:theme xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" name="Office Theme"><a:themeElements><a:clrScheme name="Office"><a:dk1><a:sysClr val="windowText" lastClr="000000"/></a:dk1><a:lt1><a:sysClr val="window" lastClr="FFFFFF"/></a:lt1><a:dk2><a:srgbClr val="44546A"/></a:dk2><a:lt2><a:srgbClr val="E7E6E6"/></a:lt2><a:accent1><a:srgbClr val="5B9BD5"/></a:accent1><a:accent2><a:srgbClr val="ED7D31"/></a:accent2><a:accent3><a:srgbClr val="A5A5A5"/></a:accent3><a:accent4><a:srgbClr val="FFC000"/></a:accent4><a:accent5><a:srgbClr val="4472C4"/></a:accent5><a:accent6><a:srgbClr val="70AD47"/></a:accent6><a:hlink><a:srgbClr val="0563C1"/></a:hlink><a:folHlink><a:srgbClr val="954F72"/></a:folHlink></a:clrScheme><a:fontScheme name="Office"><a:majorFont><a:latin typeface="Calibri Light" panose="020F0302020204030204"/><a:ea typeface=""/><a:cs typeface=""/><a:font script="Jpan" typeface="游ゴシック Light"/><a:font script="Hang" typeface="맑은 고딕"/><a:font script="Hans" typeface="等线 Light"/><a:font script="Hant" typeface="新細明體"/><a:font script="Arab" typeface="Times New Roman"/><a:font script="Hebr" typeface="Times New Roman"/><a:font script="Thai" typeface="Tahoma"/><a:font script="Ethi" typeface="Nyala"/><a:font script="Beng" typeface="Vrinda"/><a:font script="Gujr" typeface="Shruti"/><a:font script="Khmr" typeface="MoolBoran"/><a:font script="Knda" typeface="Tunga"/><a:font script="Guru" typeface="Raavi"/><a:font script="Cans" typeface="Euphemia"/><a:font script="Cher" typeface="Plantagenet Cherokee"/><a:font script="Yiii" typeface="Microsoft Yi Baiti"/><a:font script="Tibt" typeface="Microsoft Himalaya"/><a:font script="Thaa" typeface="MV Boli"/><a:font script="Deva" typeface="Mangal"/><a:font script="Telu" typeface="Gautami"/><a:font script="Taml" typeface="Latha"/><a:font script="Syrc" typeface="Estrangelo Edessa"/><a:font script="Orya" typeface="Kalinga"/><a:font script="Mlym" typeface="Kartika"/><a:font script="Laoo" typeface="DokChampa"/><a:font script="Sinh" typeface="Iskoola Pota"/><a:font script="Mong" typeface="Mongolian Baiti"/><a:font script="Viet" typeface="Times New Roman"/><a:font script="Uigh" typeface="Microsoft Uighur"/><a:font script="Geor" typeface="Sylfaen"/></a:majorFont><a:minorFont><a:latin typeface="Calibri" panose="020F0502020204030204"/><a:ea typeface=""/><a:cs typeface=""/><a:font script="Jpan" typeface="游ゴシック"/><a:font script="Hang" typeface="맑은 고딕"/><a:font script="Hans" typeface="等线"/><a:font script="Hant" typeface="新細明體"/><a:font script="Arab" typeface="Arial"/><a:font script="Hebr" typeface="Arial"/><a:font script="Thai" typeface="Tahoma"/><a:font script="Ethi" typeface="Nyala"/><a:font script="Beng" typeface="Vrinda"/><a:font script="Gujr" typeface="Shruti"/><a:font script="Khmr" typeface="DaunPenh"/><a:font script="Knda" typeface="Tunga"/><a:font script="Guru" typeface="Raavi"/><a:font script="Cans" typeface="Euphemia"/><a:font script="Cher" typeface="Plantagenet Cherokee"/><a:font script="Yiii" typeface="Microsoft Yi Baiti"/><a:font script="Tibt" typeface="Microsoft Himalaya"/><a:font script="Thaa" typeface="MV Boli"/><a:font script="Deva" typeface="Mangal"/><a:font script="Telu" typeface="Gautami"/><a:font script="Taml" typeface="Latha"/><a:font script="Syrc" typeface="Estrangelo Edessa"/><a:font script="Orya" typeface="Kalinga"/><a:font script="Mlym" typeface="Kartika"/><a:font script="Laoo" typeface="DokChampa"/><a:font script="Sinh" typeface="Iskoola Pota"/><a:font script="Mong" typeface="Mongolian Baiti"/><a:font script="Viet" typeface="Arial"/><a:font script="Uigh" typeface="Microsoft Uighur"/><a:font script="Geor" typeface="Sylfaen"/></a:minorFont></a:fontScheme><a:fmtScheme name="Office"><a:fillStyleLst><a:solidFill><a:schemeClr val="phClr"/></a:solidFill><a:gradFill rotWithShape="1"><a:gsLst><a:gs pos="0"><a:schemeClr val="phClr"><a:lumMod val="110000"/><a:satMod val="105000"/><a:tint val="67000"/></a:schemeClr></a:gs><a:gs pos="50000"><a:schemeClr val="phClr"><a:lumMod val="105000"/><a:satMod val="103000"/><a:tint val="73000"/></a:schemeClr></a:gs><a:gs pos="100000"><a:schemeClr val="phClr"><a:lumMod val="105000"/><a:satMod val="109000"/><a:tint val="81000"/></a:schemeClr></a:gs></a:gsLst><a:lin ang="5400000" scaled="0"/></a:gradFill><a:gradFill rotWithShape="1"><a:gsLst><a:gs pos="0"><a:schemeClr val="phClr"><a:satMod val="103000"/><a:lumMod val="102000"/><a:tint val="94000"/></a:schemeClr></a:gs><a:gs pos="50000"><a:schemeClr val="phClr"><a:satMod val="110000"/><a:lumMod val="100000"/><a:shade val="100000"/></a:schemeClr></a:gs><a:gs pos="100000"><a:schemeClr val="phClr"><a:lumMod val="99000"/><a:satMod val="120000"/><a:shade val="78000"/></a:schemeClr></a:gs></a:gsLst><a:lin ang="5400000" scaled="0"/></a:gradFill></a:fillStyleLst><a:lnStyleLst><a:ln w="6350" cap="flat" cmpd="sng" algn="ctr"><a:solidFill><a:schemeClr val="phClr"/></a:solidFill><a:prstDash val="solid"/><a:miter lim="800000"/></a:ln><a:ln w="12700" cap="flat" cmpd="sng" algn="ctr"><a:solidFill><a:schemeClr val="phClr"/></a:solidFill><a:prstDash val="solid"/><a:miter lim="800000"/></a:ln><a:ln w="19050" cap="flat" cmpd="sng" algn="ctr"><a:solidFill><a:schemeClr val="phClr"/></a:solidFill><a:prstDash val="solid"/><a:miter lim="800000"/></a:ln></a:lnStyleLst><a:effectStyleLst><a:effectStyle><a:effectLst/></a:effectStyle><a:effectStyle><a:effectLst/></a:effectStyle><a:effectStyle><a:effectLst><a:outerShdw blurRad="57150" dist="19050" dir="5400000" algn="ctr" rotWithShape="0"><a:srgbClr val="000000"><a:alpha val="63000"/></a:srgbClr></a:outerShdw></a:effectLst></a:effectStyle></a:effectStyleLst><a:bgFillStyleLst><a:solidFill><a:schemeClr val="phClr"/></a:solidFill><a:solidFill><a:schemeClr val="phClr"><a:tint val="95000"/><a:satMod val="170000"/></a:schemeClr></a:solidFill><a:gradFill rotWithShape="1"><a:gsLst><a:gs pos="0"><a:schemeClr val="phClr"><a:tint val="93000"/><a:satMod val="150000"/><a:shade val="98000"/><a:lumMod val="102000"/></a:schemeClr></a:gs><a:gs pos="50000"><a:schemeClr val="phClr"><a:tint val="98000"/><a:satMod val="130000"/><a:shade val="90000"/><a:lumMod val="103000"/></a:schemeClr></a:gs><a:gs pos="100000"><a:schemeClr val="phClr"><a:shade val="63000"/><a:satMod val="120000"/></a:schemeClr></a:gs></a:gsLst><a:lin ang="5400000" scaled="0"/></a:gradFill></a:bgFillStyleLst></a:fmtScheme></a:themeElements><a:objectDefaults/><a:extraClrSchemeLst/></a:theme>`

const templateNamespaceIDMap = ` xmlns="http://schemas.openxmlformats.org/spreadsheetml/2006/main" xmlns:ap="http://schemas.openxmlformats.org/officeDocument/2006/extended-properties" xmlns:op="http://schemas.openxmlformats.org/officeDocument/2006/custom-properties" xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" xmlns:c="http://schemas.openxmlformats.org/drawingml/2006/chart" xmlns:cdr="http://schemas.openxmlformats.org/drawingml/2006/chartDrawing" xmlns:comp="http://schemas.openxmlformats.org/drawingml/2006/compatibility" xmlns:dgm="http://schemas.openxmlformats.org/drawingml/2006/diagram" xmlns:lc="http://schemas.openxmlformats.org/drawingml/2006/lockedCanvas" xmlns:pic="http://schemas.openxmlformats.org/drawingml/2006/picture" xmlns:xdr="http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" xmlns:ds="http://schemas.openxmlformats.org/officeDocument/2006/customXml" xmlns:m="http://schemas.openxmlformats.org/officeDocument/2006/math" xmlns:x="http://schemas.openxmlformats.org/spreadsheetml/2006/main" xmlns:sl="http://schemas.openxmlformats.org/schemaLibrary/2006/main" xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006" xmlns:xne="http://schemas.microsoft.com/office/excel/2006/main" xmlns:mso="http://schemas.microsoft.com/office/2006/01/customui" xmlns:ax="http://schemas.microsoft.com/office/2006/activeX" xmlns:cppr="http://schemas.microsoft.com/office/2006/coverPageProps" xmlns:cdip="http://schemas.microsoft.com/office/2006/customDocumentInformationPanel" xmlns:ct="http://schemas.microsoft.com/office/2006/metadata/contentType" xmlns:ntns="http://schemas.microsoft.com/office/2006/metadata/customXsn" xmlns:lp="http://schemas.microsoft.com/office/2006/metadata/longProperties" xmlns:ma="http://schemas.microsoft.com/office/2006/metadata/properties/metaAttributes" xmlns:msink="http://schemas.microsoft.com/ink/2010/main" xmlns:c14="http://schemas.microsoft.com/office/drawing/2007/8/2/chart" xmlns:cdr14="http://schemas.microsoft.com/office/drawing/2010/chartDrawing" xmlns:a14="http://schemas.microsoft.com/office/drawing/2010/main" xmlns:pic14="http://schemas.microsoft.com/office/drawing/2010/picture" xmlns:x14="http://schemas.microsoft.com/office/spreadsheetml/2009/9/main" xmlns:xdr14="http://schemas.microsoft.com/office/excel/2010/spreadsheetDrawing" xmlns:x14ac="http://schemas.microsoft.com/office/spreadsheetml/2009/9/ac" xmlns:dsp="http://schemas.microsoft.com/office/drawing/2008/diagram" xmlns:mso14="http://schemas.microsoft.com/office/2009/07/customui" xmlns:dgm14="http://schemas.microsoft.com/office/drawing/2010/diagram" xmlns:x15="http://schemas.microsoft.com/office/spreadsheetml/2010/11/main" xmlns:x12ac="http://schemas.microsoft.com/office/spreadsheetml/2011/1/ac" xmlns:x15ac="http://schemas.microsoft.com/office/spreadsheetml/2010/11/ac" xmlns:xr="http://schemas.microsoft.com/office/spreadsheetml/2014/revision" xmlns:xr2="http://schemas.microsoft.com/office/spreadsheetml/2015/revision2" xmlns:xr3="http://schemas.microsoft.com/office/spreadsheetml/2016/revision3" xmlns:xr4="http://schemas.microsoft.com/office/spreadsheetml/2016/revision4" xmlns:xr5="http://schemas.microsoft.com/office/spreadsheetml/2016/revision5" xmlns:xr6="http://schemas.microsoft.com/office/spreadsheetml/2016/revision6" xmlns:xr7="http://schemas.microsoft.com/office/spreadsheetml/2016/revision7" xmlns:xr8="http://schemas.microsoft.com/office/spreadsheetml/2016/revision8" xmlns:xr9="http://schemas.microsoft.com/office/spreadsheetml/2016/revision9" xmlns:xr10="http://schemas.microsoft.com/office/spreadsheetml/2016/revision10" xmlns:xr11="http://schemas.microsoft.com/office/spreadsheetml/2016/revision11" xmlns:xr12="http://schemas.microsoft.com/office/spreadsheetml/2016/revision12" xmlns:xr13="http://schemas.microsoft.com/office/spreadsheetml/2016/revision13" xmlns:xr14="http://schemas.microsoft.com/office/spreadsheetml/2016/revision14" xmlns:xr15="http://schemas.microsoft.com/office/spreadsheetml/2016/revision15" xmlns:x16="http://schemas.microsoft.com/office/spreadsheetml/2014/11/main" xmlns:x16r2="http://schemas.microsoft.com/office/spreadsheetml/2015/02/main" mc:Ignorable="c14 cdr14 a14 pic14 x14 xdr14 x14ac dsp mso14 dgm14 x15 x12ac x15ac xr xr2 xr3 xr4 xr5 xr6 xr7 xr8 xr9 xr10 xr11 xr12 xr13 xr14 xr15 x15 x16 x16r2 mo mx mv o v" xmlns:mo="http://schemas.microsoft.com/office/mac/office/2008/main" xmlns:mx="http://schemas.microsoft.com/office/mac/excel/2008/main" xmlns:mv="urn:schemas-microsoft-com:mac:vml" xmlns:o="urn:schemas-microsoft-com:office:office" xmlns:v="urn:schemas-microsoft-com:vml" xr:uid="{00000000-0001-0000-0000-000000000000}">`