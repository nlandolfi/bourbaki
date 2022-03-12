# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make maximum_likelihood_density_selectors.pdf

maximum_likelihood_density_selectors.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > maximum_likelihood_density_selectors.tex

maximum_likelihood_density_selectors.pdf: ../../*.tex ../../trademark.pdf *.tex maximum_likelihood_density_selectors.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode maximum_likelihood_density_selectors.tex
	../../latexrun maximum_likelihood_density_selectors.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: maximum_likelihood_density_selectors.pdf
	open maximum_likelihood_density_selectors.pdf

o:
	make open

terms: maximum_likelihood_density_selectors.tex
	bbk_sheet -mode ts maximum_likelihood_density_selectors.tex

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
	rm -f maximum_likelihood_density_selectors.*
	make reset
	bbk_sheet -mode mk > Makefile