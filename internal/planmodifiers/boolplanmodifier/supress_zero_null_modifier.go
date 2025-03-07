package boolplanmodifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

var _ planmodifier.Bool = BoolSupressZeroNullModifierPlanModifier{}

type BoolSupressZeroNullModifierPlanModifier struct{}

// Description describes the plan modification in plain text formatting.
func (v BoolSupressZeroNullModifierPlanModifier) Description(_ context.Context) string {
	return "TODO: add plan modifier description"
}

// MarkdownDescription describes the plan modification in Markdown formatting.
func (v BoolSupressZeroNullModifierPlanModifier) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// Validate performs the plan modification.
func (v BoolSupressZeroNullModifierPlanModifier) PlanModifyBool(ctx context.Context, req planmodifier.BoolRequest, resp *planmodifier.BoolResponse) {
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"TODO: implement plan modifier SupressZeroNullModifier logic",
		req.Path.String()+": "+v.Description(ctx),
	)
}

func SupressZeroNullModifier() planmodifier.Bool {
	return BoolSupressZeroNullModifierPlanModifier{}
}
