# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make knapsack_problems.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

knapsack_problems.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > knapsack_problems.tex

knapsack_problems.pdf: ../../*.tex ../../trademark.pdf *.tex knapsack_problems.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode knapsack_problems.tex
	../../latexrun knapsack_problems.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: knapsack_problems.pdf
	open knapsack_problems.pdf

o:
	make open

terms: knapsack_problems.tex
	bbk sheet -mode ts knapsack_problems.tex

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
	rm -f knapsack_problems.*
	make reset
	bbk sheet -mode mk > Makefile