# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make stationary_covariance_functions.pdf

stationary_covariance_functions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > stationary_covariance_functions.tex

stationary_covariance_functions.pdf: ../../*.tex ../../trademark.pdf *.tex stationary_covariance_functions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode stationary_covariance_functions.tex
	../../latexrun stationary_covariance_functions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: stationary_covariance_functions.pdf
	open stationary_covariance_functions.pdf

o:
	make open

terms: stationary_covariance_functions.tex
	bbk_sheet -mode ts stationary_covariance_functions.tex

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
	rm -f stationary_covariance_functions.*
	make reset
	bbk_sheet -mode mk > Makefile