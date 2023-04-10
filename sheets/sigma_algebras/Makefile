# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make sigma_algebras.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

sigma_algebras.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > sigma_algebras.tex

sigma_algebras.pdf: ../../*.tex ../../trademark.pdf *.tex sigma_algebras.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode sigma_algebras.tex
	../../latexrun sigma_algebras.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: sigma_algebras.pdf
	open sigma_algebras.pdf

o:
	make open

terms: sigma_algebras.tex
	bbk sheet -mode ts sigma_algebras.tex

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
	rm -f sigma_algebras.*
	make reset
	bbk sheet -mode mk > Makefile