# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make tree_density_approximation_solution.pdf
	make graph.pdf

tree_density_approximation_solution.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > tree_density_approximation_solution.tex

tree_density_approximation_solution.pdf: ../../*.tex ../../trademark.pdf *.tex tree_density_approximation_solution.tex
	# pdflatex --file-line-error -interaction=nonstopmode tree_density_approximation_solution.tex
	../../latexrun tree_density_approximation_solution.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: tree_density_approximation_solution.pdf
	open tree_density_approximation_solution.pdf

o:
	make open

terms: tree_density_approximation_solution.tex
	bbk_sheet -mode ts tree_density_approximation_solution.tex

spell:
	aspell -c tree_density_approximation_solution.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f tree_density_approximation_solution.*
	make reset
	bbk_sheet -mode mk > Makefile