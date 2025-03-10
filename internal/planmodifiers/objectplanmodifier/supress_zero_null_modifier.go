package objectplanmodifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

var _ planmodifier.Object = ObjectSupressZeroNullModifierPlanModifier{}

type ObjectSupressZeroNullModifierPlanModifier struct{}

// Description describes the plan modification in plain text formatting.
func (v ObjectSupressZeroNullModifierPlanModifier) Description(_ context.Context) string {
	return "TODO: add plan modifier description"
}

// MarkdownDescription describes the plan modification in Markdown formatting.
func (v ObjectSupressZeroNullModifierPlanModifier) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// Validate performs the plan modification.
func (v ObjectSupressZeroNullModifierPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"TODO: implement plan modifier SupressZeroNullModifier logic",
		req.Path.String()+": "+v.Description(ctx),
	)
}

func SupressZeroNullModifier() planmodifier.Object {
	return ObjectSupressZeroNullModifierPlanModifier{}
}
