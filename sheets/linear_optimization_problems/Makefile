# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make linear_optimization_problems.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

linear_optimization_problems.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > linear_optimization_problems.tex

linear_optimization_problems.pdf: ../../*.tex ../../trademark.pdf *.tex linear_optimization_problems.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode linear_optimization_problems.tex
	../../latexrun linear_optimization_problems.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: linear_optimization_problems.pdf
	open linear_optimization_problems.pdf

o:
	make open

terms: linear_optimization_problems.tex
	bbk_sheet -mode ts linear_optimization_problems.tex

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
	rm -f linear_optimization_problems.*
	make reset
	bbk_sheet -mode mk > Makefile