# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make multivariate_polynomials.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

multivariate_polynomials.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > multivariate_polynomials.tex

multivariate_polynomials.pdf: ../../*.tex ../../trademark.pdf *.tex multivariate_polynomials.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode multivariate_polynomials.tex
	../../latexrun multivariate_polynomials.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: multivariate_polynomials.pdf
	open multivariate_polynomials.pdf

o:
	make open

terms: multivariate_polynomials.tex
	bbk_sheet -mode ts multivariate_polynomials.tex

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
	rm -f multivariate_polynomials.*
	make reset
	bbk_sheet -mode mk > Makefile