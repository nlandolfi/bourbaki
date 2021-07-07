# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make bounded_linear_continuous.pdf

bounded_linear_continuous.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > bounded_linear_continuous.tex

bounded_linear_continuous.pdf: ../../*.tex ../../trademark.pdf *.tex bounded_linear_continuous.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode bounded_linear_continuous.tex
	../../latexrun bounded_linear_continuous.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: bounded_linear_continuous.pdf
	open bounded_linear_continuous.pdf

o:
	make open

terms: bounded_linear_continuous.tex
	bbk_sheet -mode ts bounded_linear_continuous.tex

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
	rm -f bounded_linear_continuous.*
	make reset
	bbk_sheet -mode mk > Makefile