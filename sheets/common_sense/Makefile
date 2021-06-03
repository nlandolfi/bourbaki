# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make common_sense.pdf

common_sense.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > common_sense.tex

common_sense.pdf: ../../*.tex ../../trademark.pdf *.tex common_sense.tex
	# pdflatex --file-line-error -interaction=nonstopmode common_sense.tex
	../../latexrun common_sense.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: common_sense.pdf
	open common_sense.pdf

o:
	make open

terms: common_sense.tex
	bbk_sheet -mode ts common_sense.tex

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
	rm -f common_sense.*
	make reset
	bbk_sheet -mode mk > Makefile