# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make monotone_algebras.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

monotone_algebras.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > monotone_algebras.tex

monotone_algebras.pdf: ../../*.tex ../../trademark.pdf *.tex monotone_algebras.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode monotone_algebras.tex
	../../latexrun monotone_algebras.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: monotone_algebras.pdf
	open monotone_algebras.pdf

o:
	make open

terms: monotone_algebras.tex
	bbk_sheet -mode ts monotone_algebras.tex

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
	rm -f monotone_algebras.*
	make reset
	bbk_sheet -mode mk > Makefile