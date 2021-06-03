# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make multivariate_real_densities.pdf

multivariate_real_densities.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > multivariate_real_densities.tex

multivariate_real_densities.pdf: ../../*.tex ../../trademark.pdf *.tex multivariate_real_densities.tex
	# pdflatex --file-line-error -interaction=nonstopmode multivariate_real_densities.tex
	../../latexrun multivariate_real_densities.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: multivariate_real_densities.pdf
	open multivariate_real_densities.pdf

o:
	make open

terms: multivariate_real_densities.tex
	bbk_sheet -mode ts multivariate_real_densities.tex

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
	rm -f multivariate_real_densities.*
	make reset
	bbk_sheet -mode mk > Makefile