# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make linear_optimization_solutions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

linear_optimization_solutions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > linear_optimization_solutions.tex

linear_optimization_solutions.pdf: ../../*.tex ../../trademark.pdf *.tex linear_optimization_solutions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode linear_optimization_solutions.tex
	../../latexrun linear_optimization_solutions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: linear_optimization_solutions.pdf
	open linear_optimization_solutions.pdf

o:
	make open

terms: linear_optimization_solutions.tex
	bbk_sheet -mode ts linear_optimization_solutions.tex

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
	rm -f linear_optimization_solutions.*
	make reset
	bbk_sheet -mode mk > Makefile