# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make best_tree_distribution_approximators.pdf
	make graph.pdf

best_tree_distribution_approximators.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > best_tree_distribution_approximators.tex

best_tree_distribution_approximators.pdf: ../../*.tex ../../trademark.pdf *.tex best_tree_distribution_approximators.tex
	# pdflatex --file-line-error -interaction=nonstopmode best_tree_distribution_approximators.tex
	../../latexrun best_tree_distribution_approximators.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: best_tree_distribution_approximators.pdf
	open best_tree_distribution_approximators.pdf

o:
	make open

terms: best_tree_distribution_approximators.tex
	bbk_sheet -mode ts best_tree_distribution_approximators.tex

spell:
	aspell -c best_tree_distribution_approximators.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f best_tree_distribution_approximators.*
	make reset
	bbk_sheet -mode mk > Makefile