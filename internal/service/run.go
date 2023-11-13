package service

import (
	"fmt"
	"github.com/Nav1Cr0ss/algorithms-lab5/internal/app"
	"github.com/Nav1Cr0ss/algorithms-lab5/internal/domain"
)

func Run(populationSize, generations, knapsackCapacity, itemCount int) {
	alg := app.NewGenetic()

	items := alg.CreateItems(itemCount)

	bp := domain.NewBackpack(
		knapsackCapacity,
		items,
	)

	population := bp.InitializePopulation(populationSize)

	bestOverallFitness, bestOverallGenes := alg.CalculateGenerations(population, bp, generations)

	fmt.Println("Best BackPack Through Generations:")
	totalWeight := 0

	for i, gene := range bestOverallGenes {
		if gene {
			item := items[i]
			totalWeight += item.Weight
		}
	}
	fmt.Printf("Total Value: %d\n", bestOverallFitness)
	fmt.Printf("Total Weight: %d\n", totalWeight)
}
