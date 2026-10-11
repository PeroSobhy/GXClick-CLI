package ui

import (
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/PeroSobhy/Portfolio-CLI/internal/styles"
)

type sectionItem struct {
	key         string
	title       string
	subtitle    string
	description string
	category    string
}

func (i sectionItem) Title() string       { return i.title }
func (i sectionItem) Description() string { return i.subtitle }
func (i sectionItem) FilterValue() string { return i.title }

type menuModel struct {
	list     list.Model
	width    int
	height   int
	quantity int
}

func NewMenuModel() tea.Model {
	items := []list.Item{

		sectionItem{
			key:         "Home",
			title:       "GxClick",
			subtitle:    "You Dream, GXClick Creates",
			description: "At GXClick, we are passionate about technology and love building software that solves everyday problems in your business and personal life. Our goal is to create innovative, reliable, and user-friendly solutions that make your work easier and your life better",
			category:    "Let's building software",
		},

		sectionItem{
			key:      "Service",
			title:    "Service",
			subtitle: "Our Service",
			description: "At GXClick, we turn complex challenges into sleek, reliable digital products.\n\n" +
				"Explore our core engineering services tailored for modern businesses:\n" +
				"• Web Apps\n" +
				"• Mobile Apps\n" +
				"• Developer Tools\n" +
				"• Browser Extensions\n" +
				"• SaaS Solutions",
			category: "Portfolio",
		},

		sectionItem{
			key:      "About",
			title:    "About",
			subtitle: "About GXClick",
			description: "At GXClick, we believe technology should empower, simplify, and elevate.\n\n" +
				"We are a dedicated software venture driven by curiosity, engineering excellence, and a passion for turning bold ideas into seamless digital experiences.\n\n" +
				"Our Mission:\n" +
				"• Design, develop, and deliver high-performance software\n" +
				"• Build modern web platforms and mobile applications\n" +
				"• Create developer utilities that solve real-world problems and help creators and businesses thrive.\n\n" +
				"Our Vision:\n" +
				"To become a trusted standard for creative, scalable, and developer-friendly software solutions—turning everyday complexity into elegant, one-click reality.",
			category: "Portfolio",
		},

		sectionItem{
			key:   "Blog",
			title: "Blog",
			subtitle: "Stay up to date with our latest news, product launches, new features, and software\n" +
				"development insights. Explore tech tutorials, industry trends, and expert opinions from\n" +
				"our team to discover new ideas and learn practical skills.",
			description: "Check Our Blogs\n\nLink: [https://gxclick.vercel.app/blog]\n",
			category:    "Portfolio",
		},
	}

	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = false
	delegate.Styles.SelectedTitle = lipgloss.NewStyle().
		Background(lipgloss.Color(styles.ColorPrimary)).
		Foreground(lipgloss.Color(styles.ColorTextOnSel)).
		Padding(0, 1)
	delegate.Styles.NormalTitle = lipgloss.NewStyle().
		Foreground(lipgloss.Color(styles.ColorMuted)).
		Padding(0, 1)

	l := list.New(items, delegate, 0, 0)
	l.SetShowTitle(false)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)
	l.Select(0)

	return menuModel{
		list: l,
	}
}

func (m menuModel) Init() tea.Cmd {
	return nil
}

func (m menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "+":
			m.quantity++
		case "-":
			if m.quantity > 0 {
				m.quantity--
			}
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m menuModel) View() string {
	if m.width == 0 || m.height == 0 {
		return "Loading..."
	}

	topBar := m.renderTopBar()
	tabs := m.renderTabs()
	body := m.renderBody()

	footer := m.renderFooter()

	content := lipgloss.JoinVertical(lipgloss.Left,
		topBar,
		tabs,
		body,
		footer,
	)

	return lipgloss.NewStyle().Padding(1, 2).Render(content)
}

func (m menuModel) renderTopBar() string {
	left := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#26a269")).
		Padding(0, 1).
		Render("GxClick")

	leftSide := lipgloss.JoinHorizontal(lipgloss.Top, left, " ")

	rightSide := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#26a269")).
		Bold(true).
		Render("your favourite tech company")

	gapWidth := m.width - lipgloss.Width(leftSide) - lipgloss.Width(rightSide) - 4
	if gapWidth < 0 {
		gapWidth = 0
	}
	gap := lipgloss.NewStyle().Width(gapWidth).Render("")

	return lipgloss.JoinHorizontal(lipgloss.Top, leftSide, gap, rightSide)
}

func (m menuModel) renderTabs() string {
	thinBorder := lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "┌",
		TopRight:    "┐",
		BottomLeft:  "└",
		BottomRight: "┘",
	}

	tabStyle := lipgloss.NewStyle().
		Border(thinBorder, false, true, false, true).
		BorderForeground(lipgloss.Color(styles.ColorTextOnSel)).
		Padding(0, 1).
		Align(lipgloss.Center)

	activeTabStyle := tabStyle.Copy().
		Foreground(lipgloss.Color(styles.ColorForeground)).
		Bold(true)

	inactiveTabStyle := tabStyle.Copy().
		Foreground(lipgloss.Color(styles.ColorForeground))

	tabs := lipgloss.JoinHorizontal(lipgloss.Top,
		activeTabStyle.Render("Home"),
		inactiveTabStyle.Render("Service"),
		inactiveTabStyle.Render("About"),
		inactiveTabStyle.Render("Blogs"),
	)

	return lipgloss.NewStyle().
		Border(thinBorder, false, false, true, false).
		BorderForeground(lipgloss.Color(styles.ColorTextOnSel)).
		Width(m.width - 12).
		Align(lipgloss.Center).
		Render(tabs)
}

func (m menuModel) renderBody() string {
	sidebarWidth := 30
	contentWidth := m.width - sidebarWidth - 4

	sidebar := m.renderSidebar(sidebarWidth)
	content := m.renderContent(contentWidth)

	bodyHeight := m.height - 12
	if bodyHeight < 10 {
		bodyHeight = 10
	}

	return lipgloss.NewStyle().Height(bodyHeight).Render(lipgloss.JoinHorizontal(lipgloss.Top, sidebar, content))
}

func (m menuModel) renderSidebar(width int) string {
	var lines []string

	// Featured Header
	lines = append(lines, "")
	lines = append(lines, lipgloss.NewStyle().Foreground(lipgloss.Color("#26a269")).Render(" ~ Hello, We are ~"))

	// Featured items
	for i, item := range m.list.Items() {
		it := item.(sectionItem)
		if it.category == "Let's building software" {
			lines = append(lines, m.renderMenuItem(it, i == m.list.Index()))
		}
	}

	// Portfolio Header
	lines = append(lines, "")

	// Portfolio items
	for i, item := range m.list.Items() {
		it := item.(sectionItem)
		if it.category == "Portfolio" {
			lines = append(lines, m.renderMenuItem(it, i == m.list.Index()))
		}
	}

	return lipgloss.NewStyle().Width(width).Padding(1, 2).Render(strings.Join(lines, "\n"))
}

func (m menuModel) renderMenuItem(it sectionItem, selected bool) string {
	if selected {
		return lipgloss.NewStyle().
			Background(lipgloss.Color("#26a269")).
			Foreground(lipgloss.Color(styles.ColorTextOnSel)).
			Padding(0, 1).
			Render(it.title)
	}
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Padding(0, 1).
		Render(it.title)
}

func (m menuModel) renderContent(width int) string {
	item, ok := m.list.SelectedItem().(sectionItem)
	if !ok {
		return ""
	}

	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color(styles.ColorForeground)).
		Bold(true).
		Render(item.title)

	subtitle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(styles.ColorTextOnSel)).
		Render(item.subtitle)

	description := lipgloss.NewStyle().
		Foreground(lipgloss.Color(styles.ColorTextOnSel)).
		Width(width - 10).
		Render(item.description)

	content := lipgloss.JoinVertical(lipgloss.Left,
		title,
		"",
		subtitle,
		"",
		description,
		"",
	)

	return lipgloss.NewStyle().Padding(2, 4).Render(content)
}

func (m menuModel) renderFooter() string {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color(styles.ColorMuted))
	thinBorder := lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "┌",
		TopRight:    "┐",
		BottomLeft:  "└",
		BottomRight: "┘",
	}

	Portfolio := style.Render("↑/↓ Portfolio")
	quit := style.Render("q quit")

	footer := lipgloss.JoinHorizontal(lipgloss.Top,
		Portfolio, "   ", quit,
	)

	return lipgloss.NewStyle().
		Border(thinBorder, true, false, false, false).
		BorderForeground(lipgloss.Color(styles.ColorTextOnSel)).
		Width(m.width - 6).
		Align(lipgloss.Center).
		Render(footer)
}
