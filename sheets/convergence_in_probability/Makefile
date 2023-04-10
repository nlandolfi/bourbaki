# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make convergence_in_probability.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

convergence_in_probability.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > convergence_in_probability.tex

convergence_in_probability.pdf: ../../*.tex ../../trademark.pdf *.tex convergence_in_probability.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode convergence_in_probability.tex
	../../latexrun convergence_in_probability.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: convergence_in_probability.pdf
	open convergence_in_probability.pdf

o:
	make open

terms: convergence_in_probability.tex
	bbk sheet -mode ts convergence_in_probability.tex

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
	rm -f convergence_in_probability.*
	make reset
	bbk sheet -mode mk > Makefile