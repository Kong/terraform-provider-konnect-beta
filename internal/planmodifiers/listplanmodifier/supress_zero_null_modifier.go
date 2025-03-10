package listplanmodifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

var _ planmodifier.List = ListSupressZeroNullModifierPlanModifier{}

type ListSupressZeroNullModifierPlanModifier struct{}

// Description describes the plan modification in plain text formatting.
func (v ListSupressZeroNullModifierPlanModifier) Description(_ context.Context) string {
	return "TODO: add plan modifier description"
}

// MarkdownDescription describes the plan modification in Markdown formatting.
func (v ListSupressZeroNullModifierPlanModifier) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// Validate performs the plan modification.
func (v ListSupressZeroNullModifierPlanModifier) PlanModifyList(ctx context.Context, req planmodifier.ListRequest, resp *planmodifier.ListResponse) {
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"TODO: implement plan modifier SupressZeroNullModifier logic",
		req.Path.String()+": "+v.Description(ctx),
	)
}

func SupressZeroNullModifier() planmodifier.List {
	return ListSupressZeroNullModifierPlanModifier{}
}
