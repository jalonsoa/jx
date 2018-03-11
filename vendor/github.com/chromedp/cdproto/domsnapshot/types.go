package domsnapshot

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/domdebugger"
)

// DOMNode a Node in the DOM tree.
type DOMNode struct {
	NodeType              cdp.NodeType                 `json:"nodeType"`                        // Node's nodeType.
	NodeName              string                       `json:"nodeName"`                        // Node's nodeName.
	NodeValue             string                       `json:"nodeValue"`                       // Node's nodeValue.
	TextValue             string                       `json:"textValue,omitempty"`             // Only set for textarea elements, contains the text value.
	InputValue            string                       `json:"inputValue,omitempty"`            // Only set for input elements, contains the input's associated text value.
	InputChecked          bool                         `json:"inputChecked,omitempty"`          // Only set for radio and checkbox input elements, indicates if the element has been checked
	OptionSelected        bool                         `json:"optionSelected,omitempty"`        // Only set for option elements, indicates if the element has been selected
	BackendNodeID         cdp.BackendNodeID            `json:"backendNodeId"`                   // Node's id, corresponds to DOM.Node.backendNodeId.
	ChildNodeIndexes      []int64                      `json:"childNodeIndexes,omitempty"`      // The indexes of the node's child nodes in the domNodes array returned by getSnapshot, if any.
	Attributes            []*NameValue                 `json:"attributes,omitempty"`            // Attributes of an Element node.
	PseudoElementIndexes  []int64                      `json:"pseudoElementIndexes,omitempty"`  // Indexes of pseudo elements associated with this node in the domNodes array returned by getSnapshot, if any.
	LayoutNodeIndex       int64                        `json:"layoutNodeIndex,omitempty"`       // The index of the node's related layout tree node in the layoutTreeNodes array returned by getSnapshot, if any.
	DocumentURL           string                       `json:"documentURL,omitempty"`           // Document URL that Document or FrameOwner node points to.
	BaseURL               string                       `json:"baseURL,omitempty"`               // Base URL that Document or FrameOwner node uses for URL completion.
	ContentLanguage       string                       `json:"contentLanguage,omitempty"`       // Only set for documents, contains the document's content language.
	DocumentEncoding      string                       `json:"documentEncoding,omitempty"`      // Only set for documents, contains the document's character set encoding.
	PublicID              string                       `json:"publicId,omitempty"`              // DocumentType node's publicId.
	SystemID              string                       `json:"systemId,omitempty"`              // DocumentType node's systemId.
	FrameID               cdp.FrameID                  `json:"frameId,omitempty"`               // Frame ID for frame owner elements and also for the document node.
	ContentDocumentIndex  int64                        `json:"contentDocumentIndex,omitempty"`  // The index of a frame owner element's content document in the domNodes array returned by getSnapshot, if any.
	ImportedDocumentIndex int64                        `json:"importedDocumentIndex,omitempty"` // Index of the imported document's node of a link element in the domNodes array returned by getSnapshot, if any.
	TemplateContentIndex  int64                        `json:"templateContentIndex,omitempty"`  // Index of the content node of a template element in the domNodes array returned by getSnapshot.
	PseudoType            cdp.PseudoType               `json:"pseudoType,omitempty"`            // Type of a pseudo element node.
	ShadowRootType        cdp.ShadowRootType           `json:"shadowRootType,omitempty"`        // Shadow root type.
	IsClickable           bool                         `json:"isClickable,omitempty"`           // Whether this DOM node responds to mouse clicks. This includes nodes that have had click event listeners attached via JavaScript as well as anchor tags that naturally navigate when clicked.
	EventListeners        []*domdebugger.EventListener `json:"eventListeners,omitempty"`        // Details of the node's event listeners, if any.
	CurrentSourceURL      string                       `json:"currentSourceURL,omitempty"`      // The selected url for nodes with a srcset attribute.
}

// InlineTextBox details of post layout rendered text positions. The exact
// layout should not be regarded as stable and may change between versions.
type InlineTextBox struct {
	BoundingBox         *dom.Rect `json:"boundingBox"`         // The absolute position bounding box.
	StartCharacterIndex int64     `json:"startCharacterIndex"` // The starting index in characters, for this post layout textbox substring. Characters that would be represented as a surrogate pair in UTF-16 have length 2.
	NumCharacters       int64     `json:"numCharacters"`       // The number of characters in this post layout textbox substring. Characters that would be represented as a surrogate pair in UTF-16 have length 2.
}

// LayoutTreeNode details of an element in the DOM tree with a LayoutObject.
type LayoutTreeNode struct {
	DomNodeIndex    int64            `json:"domNodeIndex"`              // The index of the related DOM node in the domNodes array returned by getSnapshot.
	BoundingBox     *dom.Rect        `json:"boundingBox"`               // The absolute position bounding box.
	LayoutText      string           `json:"layoutText,omitempty"`      // Contents of the LayoutText, if any.
	InlineTextNodes []*InlineTextBox `json:"inlineTextNodes,omitempty"` // The post-layout inline text nodes, if any.
	StyleIndex      int64            `json:"styleIndex,omitempty"`      // Index into the computedStyles array returned by getSnapshot.
}

// ComputedStyle a subset of the full ComputedStyle as defined by the request
// whitelist.
type ComputedStyle struct {
	Properties []*NameValue `json:"properties"` // Name/value pairs of computed style properties.
}

// NameValue a name/value pair.
type NameValue struct {
	Name  string `json:"name"`  // Attribute/property name.
	Value string `json:"value"` // Attribute/property value.
}
