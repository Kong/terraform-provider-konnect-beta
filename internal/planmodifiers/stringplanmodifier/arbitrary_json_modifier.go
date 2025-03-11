package stringplanmodifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

var _ planmodifier.String = StringArbitraryJSONModifierPlanModifier{}

type StringArbitraryJSONModifierPlanModifier struct{}

// Description describes the plan modification in plain text formatting.
func (v StringArbitraryJSONModifierPlanModifier) Description(_ context.Context) string {
	return "TODO: add plan modifier description"
}

// MarkdownDescription describes the plan modification in Markdown formatting.
func (v StringArbitraryJSONModifierPlanModifier) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// Validate performs the plan modification.
func (v StringArbitraryJSONModifierPlanModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"TODO: implement plan modifier ArbitraryJSONModifier logic",
		req.Path.String()+": "+v.Description(ctx),
	)
}

func ArbitraryJSONModifier() planmodifier.String {
	return StringArbitraryJSONModifierPlanModifier{}
}
