package app

import (
	"fmt"
	"github.com/Nav1Cr0ss/algorithms-lab5/internal/domain"
	"math/rand"
	"sort"
	"time"
)

var (
	randSource    = rand.NewSource(time.Now().UnixNano())
	randGenerator = rand.New(randSource)
)

type Genetic struct {
	title       string
	generations int
}

func NewGenetic() *Genetic {

	return &Genetic{title: "Genetic Algorithm"}
}

func (g *Genetic) CreateItems(itemCount int) []domain.Item {
	items := make([]domain.Item, itemCount)

	for i := 0; i < itemCount; i++ {
		items[i] = domain.NewRandItem()
	}

	return items
}
func (g *Genetic) localImprovement(bestIndividual domain.Individual, bp domain.Backpack) domain.Individual {
	for i := range bestIndividual.Genes {
		if bestIndividual.Genes[i] {
			newGenes := make([]bool, len(bestIndividual.Genes))
			copy(newGenes, bestIndividual.Genes)
			newGenes[i] = false
			newFitness := bp.EvaluateFitness(domain.Individual{Genes: newGenes})
			if newFitness > bestIndividual.Fitness {
				bestIndividual.Genes = newGenes
				bestIndividual.Fitness = newFitness
			}
		}
	}
	return bestIndividual
}
func (g *Genetic) CalculateGenerations(population []domain.Individual, bp domain.Backpack, generations int) (int, []bool) {

	populationSize := len(population)

	bestOverallFitness := 0
	var bestOverallGenes []bool

	for generation := 0; generation < generations; generation++ {
		for i := range population {
			population[i].Fitness = bp.EvaluateFitness(population[i])
		}

		sort.SliceStable(population, func(i, j int) bool {
			return population[i].Fitness > population[j].Fitness
		})

		bestIndividual := g.localImprovement(population[0], bp)
		population[len(population)-1] = bestIndividual

		fmt.Printf("Generation %d: Best Result = %d\n", generation, bestIndividual.Fitness)

		if bestIndividual.Fitness > bestOverallFitness {
			bestOverallFitness = bestIndividual.Fitness
			bestOverallGenes = bestIndividual.Genes
		}

		newPopulation := make([]domain.Individual, populationSize)
		for i := 0; i < populationSize; i++ {
			parent1 := population[randGenerator.Intn(populationSize)]
			parent2 := population[randGenerator.Intn(populationSize)]
			childGenes := make([]bool, len(parent1.Genes))

			crossoverPoint := randGenerator.Intn(len(parent1.Genes))
			for j := 0; j < len(parent1.Genes); j++ {
				if j < crossoverPoint {
					childGenes[j] = parent1.Genes[j]
				} else {
					childGenes[j] = parent2.Genes[j]
				}
			}

			newPopulation[i] = bp.Mutate(domain.Individual{Genes: childGenes})
		}

		population = newPopulation
	}
	return bestOverallFitness, bestOverallGenes
}
