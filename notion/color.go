package notion

type AnnotationColor string

const (
	Blue             AnnotationColor = "blue"
	BlueBackground   AnnotationColor = "blue_background"
	Brown            AnnotationColor = "brown"
	BrownBackground  AnnotationColor = "brown_background"
	Default          AnnotationColor = "default"
	Gray             AnnotationColor = "gray"
	GrayBackground   AnnotationColor = "gray_background"
	Green            AnnotationColor = "green"
	GreenBackground  AnnotationColor = "green_background"
	Orange           AnnotationColor = "orange"
	OrangeBackground AnnotationColor = "orange_background"
	Pink             AnnotationColor = "pink"
	Pink_background  AnnotationColor = "pink_background"
	Purple           AnnotationColor = "purple"
	PurpleBackground AnnotationColor = "purple_background"
	Red              AnnotationColor = "red"
	RedBackground    AnnotationColor = "red_background"
	Yellow           AnnotationColor = "yellow"
	YellowBackground AnnotationColor = "yellow_background"
)

func (color AnnotationColor) Output() string {
	switch color {
	case Blue:
		return "color=blue"
	case BlueBackground:
		return "style=background:blue"
	case Brown:
		return "color=brown"
	case BrownBackground:
		return "style=background:brown"
	case Gray:
		return "color=gray"
	case GrayBackground:
		return "style=background:gray"
	case Green:
		return "color=green"
	case GreenBackground:
		return "style=background:green"
	case Orange:
		return "color=orange"
	case OrangeBackground:
		return "style=background:orange"
	case Pink:
		return "color=pink"
	case Pink_background:
		return "style=background:pink"
	case Purple:
		return "color=purple"
	case PurpleBackground:
		return "style=background:purple"
	case Red:
		return "color=red"
	case RedBackground:
		return "style=background:red"
	case Yellow:
		return "color=yellow"
	case YellowBackground:
		return "style=background:yellow"
	case Default:
	default:
		return ""
	}

	return ""
}
