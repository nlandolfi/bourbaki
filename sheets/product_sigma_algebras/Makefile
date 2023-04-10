# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make product_sigma_algebras.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

product_sigma_algebras.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > product_sigma_algebras.tex

product_sigma_algebras.pdf: ../../*.tex ../../trademark.pdf *.tex product_sigma_algebras.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode product_sigma_algebras.tex
	../../latexrun product_sigma_algebras.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: product_sigma_algebras.pdf
	open product_sigma_algebras.pdf

o:
	make open

terms: product_sigma_algebras.tex
	bbk_sheet -mode ts product_sigma_algebras.tex

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
	rm -f product_sigma_algebras.*
	make reset
	bbk_sheet -mode mk > Makefile