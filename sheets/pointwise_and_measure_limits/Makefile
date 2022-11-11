# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make pointwise_and_measure_limits.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

pointwise_and_measure_limits.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > pointwise_and_measure_limits.tex

pointwise_and_measure_limits.pdf: ../../*.tex ../../trademark.pdf *.tex pointwise_and_measure_limits.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode pointwise_and_measure_limits.tex
	../../latexrun pointwise_and_measure_limits.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: pointwise_and_measure_limits.pdf
	open pointwise_and_measure_limits.pdf

o:
	make open

terms: pointwise_and_measure_limits.tex
	bbk_sheet -mode ts pointwise_and_measure_limits.tex

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
	rm -f pointwise_and_measure_limits.*
	make reset
	bbk_sheet -mode mk > Makefile