# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make tree_distribution_approximators.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

tree_distribution_approximators.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > tree_distribution_approximators.tex

tree_distribution_approximators.pdf: ../../*.tex ../../trademark.pdf *.tex tree_distribution_approximators.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode tree_distribution_approximators.tex
	../../latexrun tree_distribution_approximators.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: tree_distribution_approximators.pdf
	open tree_distribution_approximators.pdf

o:
	make open

terms: tree_distribution_approximators.tex
	bbk sheet -mode ts tree_distribution_approximators.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f tree_distribution_approximators.*
	make reset
	bbk sheet -mode mk > Makefile