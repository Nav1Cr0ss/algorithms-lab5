package domain

const (
	mutationRate = 0.05
)

type Backpack struct {
	Capacity int
	Items    []Item
}

func NewBackpack(capacity int, items []Item) Backpack {
	return Backpack{Capacity: capacity, Items: items}

}

func (bp *Backpack) InitializePopulation(populationSize int) []Individual {
	population := make([]Individual, populationSize)

	for i := 0; i < populationSize; i++ {
		genes := make([]bool, len(bp.Items))
		for j := range genes {
			genes[j] = randGenerator.Float64() < 0.5
		}
		population[i] = Individual{Genes: genes}
	}

	return population
}

func (bp *Backpack) EvaluateFitness(individual Individual) int {
	totalValue := 0
	totalWeight := 0

	for i, gene := range individual.Genes {
		if gene {
			totalValue += bp.Items[i].Value
			totalWeight += bp.Items[i].Weight
		}
	}

	if totalWeight > bp.Capacity {
		return 0
	}

	return totalValue
}

func (bp *Backpack) Mutate(individual Individual) Individual {
	newGenes := make([]bool, len(individual.Genes))
	copy(newGenes, individual.Genes)

	for i := range newGenes {
		if randGenerator.Float64() < mutationRate {
			newGenes[i] = !newGenes[i]
		}
	}

	return Individual{Genes: newGenes}
}
