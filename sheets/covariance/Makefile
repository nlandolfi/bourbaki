# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make covariance.pdf

covariance.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > covariance.tex

covariance.pdf: ../../*.tex ../../trademark.pdf *.tex covariance.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode covariance.tex
	../../latexrun covariance.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: covariance.pdf
	open covariance.pdf

o:
	make open

terms: covariance.tex
	bbk_sheet -mode ts covariance.tex

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
	rm -f covariance.*
	make reset
	bbk_sheet -mode mk > Makefile