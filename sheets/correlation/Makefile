# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make correlation.pdf

correlation.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > correlation.tex

correlation.pdf: ../../*.tex ../../trademark.pdf *.tex correlation.tex
	# pdflatex --file-line-error -interaction=nonstopmode correlation.tex
	../../latexrun correlation.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: correlation.pdf
	open correlation.pdf

o:
	make open

terms: correlation.tex
	bbk_sheet -mode ts correlation.tex

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
	rm -f correlation.*
	make reset
	bbk_sheet -mode mk > Makefile