# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make optimal_tree_density_approximators.pdf
	make graph.pdf

optimal_tree_density_approximators.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > optimal_tree_density_approximators.tex

optimal_tree_density_approximators.pdf: ../../*.tex ../../trademark.pdf *.tex optimal_tree_density_approximators.tex
	# pdflatex --file-line-error -interaction=nonstopmode optimal_tree_density_approximators.tex
	../../latexrun optimal_tree_density_approximators.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: optimal_tree_density_approximators.pdf
	open optimal_tree_density_approximators.pdf

o:
	make open

terms: optimal_tree_density_approximators.tex
	bbk_sheet -mode ts optimal_tree_density_approximators.tex

spell:
	aspell -c optimal_tree_density_approximators.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f optimal_tree_density_approximators.*
	make reset
	bbk_sheet -mode mk > Makefile