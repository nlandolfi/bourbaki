# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make interval_length.pdf
	make graph.pdf

interval_length.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > interval_length.tex

interval_length.pdf: ../../*.tex ../../trademark.pdf *.tex interval_length.tex
	# pdflatex --file-line-error -interaction=nonstopmode interval_length.tex
	../../latexrun interval_length.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: interval_length.pdf
	open interval_length.pdf

o:
	make open

terms: interval_length.tex
	bbk_sheet -mode ts interval_length.tex

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
	rm -f interval_length.*
	make reset
	bbk_sheet -mode mk > Makefile