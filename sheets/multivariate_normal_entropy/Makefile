# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make multivariate_normal_entropy.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

multivariate_normal_entropy.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > multivariate_normal_entropy.tex

multivariate_normal_entropy.pdf: ../../*.tex ../../trademark.pdf *.tex multivariate_normal_entropy.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode multivariate_normal_entropy.tex
	../../latexrun multivariate_normal_entropy.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: multivariate_normal_entropy.pdf
	open multivariate_normal_entropy.pdf

o:
	make open

terms: multivariate_normal_entropy.tex
	bbk sheet -mode ts multivariate_normal_entropy.tex

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
	rm -f multivariate_normal_entropy.*
	make reset
	bbk sheet -mode mk > Makefile