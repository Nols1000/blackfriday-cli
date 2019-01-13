package main

import (
	"flag"
	"github.com/Nols1000/blackfriday-vbcode"
	"gopkg.in/russross/blackfriday.v2"
	"io/ioutil"
	"log"
	"os"
)

const VERSION  = "0.1.0"

func main() {
	var err error
	var format, footnoteAnchorPrefix, footnoteReturnLinkContent, headingIdPrefix, headingIdSuffix, title, css,
	icon string
	var headingLevelOffset int
	var skipHTML, skipImages, skipLinks, safelink, nofollowLinks, noreferrerLinks, noopenerLinks, hrefTargetBlank,
	completePage, footnoteReturnLinks, smartypants, smartypantsFractions, smartypantsDashes, smartypantsLatexDashes,
	smartypantsAngledQuotes, smartypantsQuotesNBSP, toc, commonHTMLFlags, noIntraEmphasis, tables, fencedCode, autolink,
	strikethrough, laxHTMLBlocks, spaceHeadings, hardLineBreak, tabSizeEight, footnotes, noEmptyLineBeforeBlock,
	headingIDs, titleblock, autoHeadingIDs, backslashLineBreak, definitionList, commonExtensions bool

	flag.StringVar(&format, "format", "html", "Format of the output (html, xhtml, vb-code)")
	flag.StringVar(&footnoteAnchorPrefix, "footnote-anchor-prefix", "", "")
	flag.StringVar(&footnoteReturnLinkContent, "footnote-return-link-content", "", "")
	flag.StringVar(&headingIdPrefix, "heading-id-prefix", "", "")
	flag.StringVar(&headingIdSuffix, "heading-id-suffix", "", "")
	flag.StringVar(&title, "title", "", "")
	flag.StringVar(&css, "css", "", "")
	flag.StringVar(&icon, "icon", "", "")

	flag.IntVar(&headingLevelOffset, "heading-level-offset", 0, "")

	flag.BoolVar(&skipHTML, "skip-html", false, "")
	flag.BoolVar(&skipImages, "skip-images", false, "")
	flag.BoolVar(&skipLinks, "skip-links", false, "")
	flag.BoolVar(&safelink, "safelink", false, "")
	flag.BoolVar(&noreferrerLinks, "noreferrer-links", false, "")
	flag.BoolVar(&nofollowLinks, "nofollow-links", false, "")
	flag.BoolVar(&noopenerLinks, "noopener-links", false, "")
	flag.BoolVar(&hrefTargetBlank, "href-target-blank", false, "")
	flag.BoolVar(&completePage, "complete-page", false, "")
	flag.BoolVar(&footnoteReturnLinks, "footnote-return-links", false, "")
	flag.BoolVar(&smartypants, "smartypants", false, "")
	flag.BoolVar(&smartypantsFractions, "smartypants-fractions", false, "")
	flag.BoolVar(&smartypantsDashes, "smartypants-dashes", false, "")
	flag.BoolVar(&smartypantsLatexDashes, "smartypants-latex-dashes", false, "")
	flag.BoolVar(&smartypantsAngledQuotes, "smartypants-angled-quotes", false, "")
	flag.BoolVar(&smartypantsQuotesNBSP, "smartypants-quotes-nbsp", false, "")
	flag.BoolVar(&toc, "toc", false, "")
	flag.BoolVar(&commonHTMLFlags, "common-html-flags", false, "")

	flag.BoolVar(&noIntraEmphasis, "no-intra-emphasis", false, "")
	flag.BoolVar(&tables, "tables", false, "")
	flag.BoolVar(&fencedCode, "fenced-code", false, "")
	flag.BoolVar(&autolink, "autolink", false, "")
	flag.BoolVar(&strikethrough, "strikethrough", false, "")
	flag.BoolVar(&laxHTMLBlocks, "lax-html-blocks", false, "")
	flag.BoolVar(&spaceHeadings, "space-headings", false, "")
	flag.BoolVar(&hardLineBreak, "hard-line-break", false, "")
	flag.BoolVar(&tabSizeEight, "tab-size-eight", false, "")
	flag.BoolVar(&footnotes, "footnotes", false, "")
	flag.BoolVar(&noEmptyLineBeforeBlock, "no-empty-line-before-blocks", false, "")
	flag.BoolVar(&headingIDs, "heading-ids", false, "")
	flag.BoolVar(&titleblock, "titleblock", false, "")
	flag.BoolVar(&autoHeadingIDs, "auto-heading-ids", false, "")
	flag.BoolVar(&backslashLineBreak, "backslash-line-break", false, "")
	flag.BoolVar(&definitionList, "definition-list", false, "")
	flag.BoolVar(&commonExtensions, "common-extensions", false, "")

	flag.Usage = func() {
		log.Printf( "Blackfriday Markdown Processor (with blackfriday-latex and blackfriday-vbcode) v" + VERSION +
			"\nAvailable at http://github.com/Nols1000/blackfriday-cli\n\n" +
			"Blackfriday Cli and Blackfriday vBullettin Code\n" +
			"Copyright © 2019 Nils-Börge Margotti <nilsmargotti@gmail.com>\n" +
			"Distributed under the MIT License\n"+
			"Blackfriday\n" +
			"Copyright © 2011 Russ Ross <russ@russross.com>\n" +
			"Distributed under the Simplified BSD License\n"+
			"See website for details\n\n"+
			"Usage:\n"+
			"  %s [options] [inputfile [outputfile]]\n\n"+
			"Options:\n",
			os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	args := flag.Args()

	var input []byte

	switch len(args) {
	case 0:
		if input, err = ioutil.ReadAll(os.Stdin); err != nil {
			log.Fatalf( "Error reading from Stdin:", err)
			os.Exit(-1)
		}
	case 1, 2:
		if input, err = ioutil.ReadFile(args[0]); err != nil {
			log.Fatalf( "Error reading from", args[0], ":", err)
			os.Exit(-1)
		}
	default:
		flag.Usage()
		os.Exit(-1)
	}


	// HTML-Flags
	htmlFlags := blackfriday.HTMLFlagsNone

	if skipHTML {
		htmlFlags |= blackfriday.SkipHTML
	}

	if skipImages {
		htmlFlags |= blackfriday.SkipImages
	}

	if skipLinks {
		htmlFlags |= blackfriday.SkipLinks
	}

	if safelink {
		htmlFlags |= blackfriday.Safelink
	}

	if nofollowLinks {
		htmlFlags |= blackfriday.NofollowLinks
	}

	if noreferrerLinks {
		htmlFlags |= blackfriday.NoreferrerLinks
	}

	if noopenerLinks {
		htmlFlags |= blackfriday.NoopenerLinks
	}

	if hrefTargetBlank {
		htmlFlags |= blackfriday.HrefTargetBlank
	}

	if completePage {
		htmlFlags |= blackfriday.CompletePage
	}

	if footnoteReturnLinks {
		htmlFlags |= blackfriday.FootnoteReturnLinks
	}

	if smartypants {
		htmlFlags |= blackfriday.Smartypants
	}

	if smartypantsFractions {
		htmlFlags |= blackfriday.SmartypantsFractions
	}

	if smartypantsDashes {
		htmlFlags |= blackfriday.SmartypantsDashes
	}

	if smartypantsLatexDashes {
		htmlFlags |= blackfriday.SmartypantsLatexDashes
	}

	if smartypantsAngledQuotes {
		htmlFlags |= blackfriday.SmartypantsAngledQuotes
	}

	if smartypantsQuotesNBSP {
		htmlFlags |= blackfriday.SmartypantsQuotesNBSP
	}

	if toc {
		htmlFlags |= blackfriday.TOC
	}

	if commonHTMLFlags {
		htmlFlags |= blackfriday.CommonHTMLFlags
	}

	// Extensions
	extensions := blackfriday.NoExtensions

	if noIntraEmphasis {
		extensions |= blackfriday.NoIntraEmphasis
	}

	if tables {
		extensions |= blackfriday.Tables
	}

	if fencedCode {
		extensions |= blackfriday.FencedCode
	}

	if autolink {
		extensions |= blackfriday.Autolink
	}

	if strikethrough {
		extensions |= blackfriday.Strikethrough
	}

	if laxHTMLBlocks {
		extensions |= blackfriday.LaxHTMLBlocks
	}

	if spaceHeadings {
		extensions |= blackfriday.SpaceHeadings
	}

	if hardLineBreak {
		extensions |= blackfriday.HardLineBreak
	}

	if tabSizeEight {
		extensions |= blackfriday.TabSizeEight
	}

	if footnotes {
		extensions |= blackfriday.Footnotes
	}

	if noEmptyLineBeforeBlock {
		extensions |= blackfriday.NoEmptyLineBeforeBlock
	}

	if headingIDs {
		extensions |= blackfriday.HeadingIDs
	}

	if titleblock {
		extensions |= blackfriday.Titleblock
	}

	if autoHeadingIDs {
		extensions |= blackfriday.AutoHeadingIDs
	}

	if backslashLineBreak {
		extensions |= blackfriday.BackslashLineBreak
	}

	if definitionList {
		extensions |= blackfriday.DefinitionLists
	}

	if commonExtensions {
		extensions |= blackfriday.CommonExtensions
	}

	var renderer blackfriday.Renderer

	switch format {
	case "html":
		renderer = blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
			FootnoteAnchorPrefix: footnoteAnchorPrefix,
			FootnoteReturnLinkContents: footnoteReturnLinkContent,
			HeadingIDPrefix: headingIdPrefix,
			HeadingIDSuffix: headingIdSuffix,
			HeadingLevelOffset: headingLevelOffset,
			Title: title,
			CSS: css,
			Icon: icon,
			Flags: htmlFlags,
		})
	case "xhtml":
		renderer = blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
			FootnoteAnchorPrefix: footnoteAnchorPrefix,
			FootnoteReturnLinkContents: footnoteReturnLinkContent,
			HeadingIDPrefix: headingIdPrefix,
			HeadingIDSuffix: headingIdSuffix,
			HeadingLevelOffset: headingLevelOffset,
			Title: title,
			CSS: css,
			Icon: icon,
			Flags: htmlFlags | blackfriday.UseXHTML,
		})
	case "vb-code":
		renderer = blackfriday_vbcode.NewVBulletinRenderer()
	default:
		log.Fatalf("The output format %s is not supported. Try one of the supported formats html, xhtml or vb-code.", format)
		flag.Usage()
		os.Exit(-1)
	}

	var output []byte
	output = blackfriday.Run(input, blackfriday.WithRenderer(renderer), blackfriday.WithExtensions(extensions))

	// output the result
	var out *os.File
	if len(args) == 2 {
		if out, err = os.Create(args[1]); err != nil {
			log.Fatalf("Error creating %s: %v", args[1], err)
			os.Exit(-1)
		}
		defer out.Close()
	} else {
		out = os.Stdout
	}

	if _, err = out.Write(output); err != nil {
		log.Fatalf( "Error writing output: %v", err)
		os.Exit(-1)
	}
}