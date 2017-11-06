// Copyright 2014 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package declarative

import (
	"github.com/lxn/walk"
)

type ScrollView struct {
	// Window

	Background         Brush
	ContextMenuItems   []MenuItem
	Enabled            Property
	Font               Font
	MaxSize            Size
	MinSize            Size
	Name               string
	OnBoundsChanged    walk.EventHandler
	OnKeyDown          walk.KeyEventHandler
	OnKeyPress         walk.KeyEventHandler
	OnKeyUp            walk.KeyEventHandler
	OnMouseDown        walk.MouseEventHandler
	OnMouseMove        walk.MouseEventHandler
	OnMouseUp          walk.MouseEventHandler
	OnSizeChanged      walk.EventHandler
	Persistent         bool
	RightToLeftReading bool
	ToolTipText        Property
	Visible            Property

	// Widget

	AlwaysConsumeSpace bool
	Column             int
	ColumnSpan         int
	Row                int
	RowSpan            int
	StretchFactor      int

	// Container

	Children   []Widget
	DataBinder DataBinder
	Layout     Layout

	// ScrollView

	AssignTo        **walk.ScrollView
	HorizontalFixed bool
	VerticalFixed   bool
}

func (sv ScrollView) Create(builder *Builder) error {
	w, err := walk.NewScrollView(builder.Parent())
	if err != nil {
		return err
	}

	w.SetSuspended(true)
	builder.Defer(func() error {
		w.SetSuspended(false)
		return nil
	})

	w.SetScrollbars(!sv.HorizontalFixed, !sv.VerticalFixed)

	return builder.InitWidget(sv, w, func() error {
		if sv.AssignTo != nil {
			*sv.AssignTo = w
		}

		return nil
	})
}