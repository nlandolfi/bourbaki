# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make pAC_learnable_hypotheses.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

pAC_learnable_hypotheses.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > pAC_learnable_hypotheses.tex

pAC_learnable_hypotheses.pdf: ../../*.tex ../../trademark.pdf *.tex pAC_learnable_hypotheses.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode pAC_learnable_hypotheses.tex
	../../latexrun pAC_learnable_hypotheses.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: pAC_learnable_hypotheses.pdf
	open pAC_learnable_hypotheses.pdf

o:
	make open

terms: pAC_learnable_hypotheses.tex
	bbk sheet -mode ts pAC_learnable_hypotheses.tex

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
	rm -f pAC_learnable_hypotheses.*
	make reset
	bbk sheet -mode mk > Makefile