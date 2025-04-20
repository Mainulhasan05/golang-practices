package main

import (
	"fmt"
)

type job_experience struct {
	company_name string
	months       float32
	role         string
	tech_stack   []string
	isRunning    bool
}

var name string
var experiences [2]job_experience

func init() {
	name = "Md. Mainul Hasan"
	experiences[0] = job_experience{"Microdeft", 23, "Software Engineer", []string{"React", "Next.js", "Python", "Flutter", "Node.js", "Selenium Automation"}, false}
	experiences[1] = job_experience{"Roshoon", 5, "Backend Dev + Team Lead", []string{"Node.js", "PostgreSQL", "Next.js"}, true}
}

func formatDuration(months float32) string {
	years := int(months) / 12
	remainingMonths := int(months) % 12
	if years > 0 {
		return fmt.Sprintf("%d yr %d mo", years, remainingMonths)
	}
	return fmt.Sprintf("%d mo", remainingMonths)
}

func printInfo() {
	fmt.Println()
	fmt.Println("===========================================")
	fmt.Printf("👤 Name: %s\n", name)
	fmt.Println("===========================================")

	for _, exp := range experiences {
		status := "✅ Completed"
		if exp.isRunning {
			status = "🟢 Currently Working"
		}

		fmt.Printf("\n🏢 Company: %s\n", exp.company_name)
		fmt.Printf("📅 Duration: %s\n", formatDuration(exp.months))
		fmt.Printf("💼 Role: %s\n", exp.role)
		fmt.Printf("📍 Status: %s\n", status)
		fmt.Println("🛠️  Tech Stack:")
		for _, tech := range exp.tech_stack {
			fmt.Printf("   • %s\n", tech)
		}
		fmt.Println("-------------------------------------------")
	}
}

func main() {
	printInfo()
}
