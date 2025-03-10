// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect-beta/internal/provider/types"
	"github.com/kong/terraform-provider-konnect-beta/internal/sdk/models/shared"
)

func (r *PortalCustomizationResourceModel) ToSharedPortalCustomization() *shared.PortalCustomization {
	var theme *shared.Theme
	if r.Theme != nil {
		name := new(string)
		if !r.Theme.Name.IsUnknown() && !r.Theme.Name.IsNull() {
			*name = r.Theme.Name.ValueString()
		} else {
			name = nil
		}
		mode := new(shared.Mode)
		if !r.Theme.Mode.IsUnknown() && !r.Theme.Mode.IsNull() {
			*mode = shared.Mode(r.Theme.Mode.ValueString())
		} else {
			mode = nil
		}
		var colors *shared.Colors
		if r.Theme.Colors != nil {
			primary := new(string)
			if !r.Theme.Colors.Primary.IsUnknown() && !r.Theme.Colors.Primary.IsNull() {
				*primary = r.Theme.Colors.Primary.ValueString()
			} else {
				primary = nil
			}
			colors = &shared.Colors{
				Primary: primary,
			}
		}
		theme = &shared.Theme{
			Name:   name,
			Mode:   mode,
			Colors: colors,
		}
	}
	layout := new(string)
	if !r.Layout.IsUnknown() && !r.Layout.IsNull() {
		*layout = r.Layout.ValueString()
	} else {
		layout = nil
	}
	css := new(string)
	if !r.CSS.IsUnknown() && !r.CSS.IsNull() {
		*css = r.CSS.ValueString()
	} else {
		css = nil
	}
	var js *shared.Js
	if r.Js != nil {
		custom := new(string)
		if !r.Js.Custom.IsUnknown() && !r.Js.Custom.IsNull() {
			*custom = r.Js.Custom.ValueString()
		} else {
			custom = nil
		}
		var scripts []string = []string{}
		for _, scriptsItem := range r.Js.Scripts {
			scripts = append(scripts, scriptsItem.ValueString())
		}
		js = &shared.Js{
			Custom:  custom,
			Scripts: scripts,
		}
	}
	var menu *shared.Menu
	if r.Menu != nil {
		var main []shared.PortalMenuItem = []shared.PortalMenuItem{}
		for _, mainItem := range r.Menu.Main {
			var path string
			path = mainItem.Path.ValueString()

			var title string
			title = mainItem.Title.ValueString()

			visibility := shared.Visibility(mainItem.Visibility.ValueString())
			var external bool
			external = mainItem.External.ValueBool()

			main = append(main, shared.PortalMenuItem{
				Path:       path,
				Title:      title,
				Visibility: visibility,
				External:   external,
			})
		}
		var footerSections []shared.PortalFooterMenuSection = []shared.PortalFooterMenuSection{}
		for _, footerSectionsItem := range r.Menu.FooterSections {
			var title1 string
			title1 = footerSectionsItem.Title.ValueString()

			var items []shared.PortalMenuItem = []shared.PortalMenuItem{}
			for _, itemsItem := range footerSectionsItem.Items {
				var path1 string
				path1 = itemsItem.Path.ValueString()

				var title2 string
				title2 = itemsItem.Title.ValueString()

				visibility1 := shared.Visibility(itemsItem.Visibility.ValueString())
				var external1 bool
				external1 = itemsItem.External.ValueBool()

				items = append(items, shared.PortalMenuItem{
					Path:       path1,
					Title:      title2,
					Visibility: visibility1,
					External:   external1,
				})
			}
			footerSections = append(footerSections, shared.PortalFooterMenuSection{
				Title: title1,
				Items: items,
			})
		}
		var footerBottom []shared.PortalMenuItem = []shared.PortalMenuItem{}
		for _, footerBottomItem := range r.Menu.FooterBottom {
			var path2 string
			path2 = footerBottomItem.Path.ValueString()

			var title3 string
			title3 = footerBottomItem.Title.ValueString()

			visibility2 := shared.Visibility(footerBottomItem.Visibility.ValueString())
			var external2 bool
			external2 = footerBottomItem.External.ValueBool()

			footerBottom = append(footerBottom, shared.PortalMenuItem{
				Path:       path2,
				Title:      title3,
				Visibility: visibility2,
				External:   external2,
			})
		}
		menu = &shared.Menu{
			Main:           main,
			FooterSections: footerSections,
			FooterBottom:   footerBottom,
		}
	}
	var specRenderer *shared.SpecRenderer
	if r.SpecRenderer != nil {
		tryItUI := new(bool)
		if !r.SpecRenderer.TryItUI.IsUnknown() && !r.SpecRenderer.TryItUI.IsNull() {
			*tryItUI = r.SpecRenderer.TryItUI.ValueBool()
		} else {
			tryItUI = nil
		}
		tryItInsomnia := new(bool)
		if !r.SpecRenderer.TryItInsomnia.IsUnknown() && !r.SpecRenderer.TryItInsomnia.IsNull() {
			*tryItInsomnia = r.SpecRenderer.TryItInsomnia.ValueBool()
		} else {
			tryItInsomnia = nil
		}
		infiniteScroll := new(bool)
		if !r.SpecRenderer.InfiniteScroll.IsUnknown() && !r.SpecRenderer.InfiniteScroll.IsNull() {
			*infiniteScroll = r.SpecRenderer.InfiniteScroll.ValueBool()
		} else {
			infiniteScroll = nil
		}
		showSchemas := new(bool)
		if !r.SpecRenderer.ShowSchemas.IsUnknown() && !r.SpecRenderer.ShowSchemas.IsNull() {
			*showSchemas = r.SpecRenderer.ShowSchemas.ValueBool()
		} else {
			showSchemas = nil
		}
		specRenderer = &shared.SpecRenderer{
			TryItUI:        tryItUI,
			TryItInsomnia:  tryItInsomnia,
			InfiniteScroll: infiniteScroll,
			ShowSchemas:    showSchemas,
		}
	}
	robots := new(string)
	if !r.Robots.IsUnknown() && !r.Robots.IsNull() {
		*robots = r.Robots.ValueString()
	} else {
		robots = nil
	}
	out := shared.PortalCustomization{
		Theme:        theme,
		Layout:       layout,
		CSS:          css,
		Js:           js,
		Menu:         menu,
		SpecRenderer: specRenderer,
		Robots:       robots,
	}
	return &out
}

func (r *PortalCustomizationResourceModel) RefreshFromSharedPortalCustomization(resp *shared.PortalCustomization) {
	if resp != nil {
		r.CSS = types.StringPointerValue(resp.CSS)
		if resp.Js == nil {
			r.Js = nil
		} else {
			r.Js = &tfTypes.Js{}
			r.Js.Custom = types.StringPointerValue(resp.Js.Custom)
			r.Js.Scripts = make([]types.String, 0, len(resp.Js.Scripts))
			for _, v := range resp.Js.Scripts {
				r.Js.Scripts = append(r.Js.Scripts, types.StringValue(v))
			}
		}
		r.Layout = types.StringPointerValue(resp.Layout)
		if resp.Menu == nil {
			r.Menu = nil
		} else {
			r.Menu = &tfTypes.Menu{}
			r.Menu.FooterBottom = []tfTypes.PortalMenuItem{}
			if len(r.Menu.FooterBottom) > len(resp.Menu.FooterBottom) {
				r.Menu.FooterBottom = r.Menu.FooterBottom[:len(resp.Menu.FooterBottom)]
			}
			for footerBottomCount, footerBottomItem := range resp.Menu.FooterBottom {
				var footerBottom1 tfTypes.PortalMenuItem
				footerBottom1.External = types.BoolValue(footerBottomItem.External)
				footerBottom1.Path = types.StringValue(footerBottomItem.Path)
				footerBottom1.Title = types.StringValue(footerBottomItem.Title)
				footerBottom1.Visibility = types.StringValue(string(footerBottomItem.Visibility))
				if footerBottomCount+1 > len(r.Menu.FooterBottom) {
					r.Menu.FooterBottom = append(r.Menu.FooterBottom, footerBottom1)
				} else {
					r.Menu.FooterBottom[footerBottomCount].External = footerBottom1.External
					r.Menu.FooterBottom[footerBottomCount].Path = footerBottom1.Path
					r.Menu.FooterBottom[footerBottomCount].Title = footerBottom1.Title
					r.Menu.FooterBottom[footerBottomCount].Visibility = footerBottom1.Visibility
				}
			}
			r.Menu.FooterSections = []tfTypes.PortalFooterMenuSection{}
			if len(r.Menu.FooterSections) > len(resp.Menu.FooterSections) {
				r.Menu.FooterSections = r.Menu.FooterSections[:len(resp.Menu.FooterSections)]
			}
			for footerSectionsCount, footerSectionsItem := range resp.Menu.FooterSections {
				var footerSections1 tfTypes.PortalFooterMenuSection
				footerSections1.Items = []tfTypes.PortalMenuItem{}
				for itemsCount, itemsItem := range footerSectionsItem.Items {
					var items1 tfTypes.PortalMenuItem
					items1.External = types.BoolValue(itemsItem.External)
					items1.Path = types.StringValue(itemsItem.Path)
					items1.Title = types.StringValue(itemsItem.Title)
					items1.Visibility = types.StringValue(string(itemsItem.Visibility))
					if itemsCount+1 > len(footerSections1.Items) {
						footerSections1.Items = append(footerSections1.Items, items1)
					} else {
						footerSections1.Items[itemsCount].External = items1.External
						footerSections1.Items[itemsCount].Path = items1.Path
						footerSections1.Items[itemsCount].Title = items1.Title
						footerSections1.Items[itemsCount].Visibility = items1.Visibility
					}
				}
				footerSections1.Title = types.StringValue(footerSectionsItem.Title)
				if footerSectionsCount+1 > len(r.Menu.FooterSections) {
					r.Menu.FooterSections = append(r.Menu.FooterSections, footerSections1)
				} else {
					r.Menu.FooterSections[footerSectionsCount].Items = footerSections1.Items
					r.Menu.FooterSections[footerSectionsCount].Title = footerSections1.Title
				}
			}
			r.Menu.Main = []tfTypes.PortalMenuItem{}
			if len(r.Menu.Main) > len(resp.Menu.Main) {
				r.Menu.Main = r.Menu.Main[:len(resp.Menu.Main)]
			}
			for mainCount, mainItem := range resp.Menu.Main {
				var main1 tfTypes.PortalMenuItem
				main1.External = types.BoolValue(mainItem.External)
				main1.Path = types.StringValue(mainItem.Path)
				main1.Title = types.StringValue(mainItem.Title)
				main1.Visibility = types.StringValue(string(mainItem.Visibility))
				if mainCount+1 > len(r.Menu.Main) {
					r.Menu.Main = append(r.Menu.Main, main1)
				} else {
					r.Menu.Main[mainCount].External = main1.External
					r.Menu.Main[mainCount].Path = main1.Path
					r.Menu.Main[mainCount].Title = main1.Title
					r.Menu.Main[mainCount].Visibility = main1.Visibility
				}
			}
		}
		r.Robots = types.StringPointerValue(resp.Robots)
		if resp.SpecRenderer == nil {
			r.SpecRenderer = nil
		} else {
			r.SpecRenderer = &tfTypes.SpecRenderer{}
			r.SpecRenderer.InfiniteScroll = types.BoolPointerValue(resp.SpecRenderer.InfiniteScroll)
			r.SpecRenderer.ShowSchemas = types.BoolPointerValue(resp.SpecRenderer.ShowSchemas)
			r.SpecRenderer.TryItInsomnia = types.BoolPointerValue(resp.SpecRenderer.TryItInsomnia)
			r.SpecRenderer.TryItUI = types.BoolPointerValue(resp.SpecRenderer.TryItUI)
		}
		if resp.Theme == nil {
			r.Theme = nil
		} else {
			r.Theme = &tfTypes.Theme{}
			if resp.Theme.Colors == nil {
				r.Theme.Colors = nil
			} else {
				r.Theme.Colors = &tfTypes.Colors{}
				r.Theme.Colors.Primary = types.StringPointerValue(resp.Theme.Colors.Primary)
			}
			if resp.Theme.Mode != nil {
				r.Theme.Mode = types.StringValue(string(*resp.Theme.Mode))
			} else {
				r.Theme.Mode = types.StringNull()
			}
			r.Theme.Name = types.StringPointerValue(resp.Theme.Name)
		}
	}
}
