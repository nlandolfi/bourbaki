# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make tail_measure_upper_bound.pdf

tail_measure_upper_bound.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > tail_measure_upper_bound.tex

tail_measure_upper_bound.pdf: ../../*.tex ../../trademark.pdf *.tex tail_measure_upper_bound.tex
	# pdflatex --file-line-error -interaction=nonstopmode tail_measure_upper_bound.tex
	../../latexrun tail_measure_upper_bound.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: tail_measure_upper_bound.pdf
	open tail_measure_upper_bound.pdf

o:
	make open

terms: tail_measure_upper_bound.tex
	bbk_sheet -mode ts tail_measure_upper_bound.tex

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
	rm -f tail_measure_upper_bound.*
	make reset
	bbk_sheet -mode mk > Makefile