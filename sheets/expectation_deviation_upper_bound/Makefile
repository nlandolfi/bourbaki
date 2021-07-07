# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make expectation_deviation_upper_bound.pdf

expectation_deviation_upper_bound.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > expectation_deviation_upper_bound.tex

expectation_deviation_upper_bound.pdf: ../../*.tex ../../trademark.pdf *.tex expectation_deviation_upper_bound.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode expectation_deviation_upper_bound.tex
	../../latexrun expectation_deviation_upper_bound.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: expectation_deviation_upper_bound.pdf
	open expectation_deviation_upper_bound.pdf

o:
	make open

terms: expectation_deviation_upper_bound.tex
	bbk_sheet -mode ts expectation_deviation_upper_bound.tex

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
	rm -f expectation_deviation_upper_bound.*
	make reset
	bbk_sheet -mode mk > Makefile